package basic

import (
	"net/http"
	"golang.org/x/net/websocket"
	"github.com/labstack/gommon/log"
	"fmt"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Stamp time.Time `json:"stamp"`
}

func echoHandler(ws *websocket.Conn) {
	fmt.Printf("Remote Addr %s\n", ws.RemoteAddr())
	buf := make([]byte, 512)
	n, err := ws.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive msg: %s\n", buf[:n])
	msg := []byte("[" + string(buf[:n]) + "]")
	n, err = ws.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send msg: %s", msg[:n])

	ws.Write([]byte("This is send message"))

	p := &Person{
		Name: "benjamin",
		Age:  21,
	}
	websocket.JSON.Send(ws, p)
}
func GameServe() http.Handler {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	return http.DefaultServeMux

}
