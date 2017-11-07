package web

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/gorilla/websocket"
	"time"
	"github.com/labstack/gommon/log"
	"bytes"
	"fmt"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	upgrader2 = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024}
)
var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for true {

		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for true {
		select {
		case message, ok := <-c.send:
			fmt.Printf("send message %s\n", message)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			ws, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			ws.Write(message)
			n := len(c.send)
			for i := 0; i < n; i++ {
				ws.Write(newline)
				ws.Write(<-c.send)
			}
			if err := ws.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Error: %s\n", err)
			}
			break
		}
		fmt.Printf("receive message %s\n", message)
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
		fmt.Printf("receive message2 %s\n", message)
	}

}

func serveHome(ctx echo.Context) error {
	ctx.Logger().Printf("Request url: %s\n", ctx.Path())
	if ctx.Request().URL.Path != "/" {
		return ctx.String(http.StatusNotFound, "Not found")
	}
	if ctx.Request().Method != echo.GET {
		return ctx.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
	return ctx.File("web/html/home.html")
}

func serveWs(hub *Hub, ctx echo.Context) error {
	conn, err := upgrader2.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		ctx.Logger().Error(err)
		return err
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	go client.writePump()
	go client.readPump()
	return nil
}

func EchoGorilla2Serve() http.Handler {
	hub := NewHub()
	go hub.run()
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.GET("/", serveHome)
	app.GET("/ws", func(ctx echo.Context) error {
		return serveWs(hub, ctx)
	})
	return app
}
