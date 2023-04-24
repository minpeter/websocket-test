package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	fmt.Printf("Request: %+v\n", ws.Request())
	io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe("0.0.0.0:10000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
