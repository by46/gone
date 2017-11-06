package web

import (
	"github.com/labstack/echo"
	"html/template"
	"net/http"
	"github.com/labstack/echo/middleware"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
	"fmt"
)

var upgrader = websocket.Upgrader{}

func wsEcho(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}

		fmt.Printf("Receive message: %v, %s\n", mt, message)
		if err = ws.WriteMessage(mt, message); err != nil {
			log.Printf("Send message error: %v\n", err)
			break
		}

	}
	return nil
}

func home(ctx echo.Context) error {
	return homeTemplate.Execute(ctx.Response(), "ws://"+ctx.Request().Host+"/ws")
}
func EchoGorillaServe() http.Handler {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.GET("/", home)
	app.GET("/ws", wsEcho)
	return app
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server,
"Send" to send a message to the server and "Close" to close the connection.
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
