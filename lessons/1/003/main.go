// GET HTTP/1.1
// Host: echo.websocket.org
// Upgrade: websocket
// Connection: Upgrade
// Sec-Websocket-Key: jhfkgcjhdk
// Sec-Websocket-Version: 13

// RS:
// HTTP/1.1 101 Switching Protocols
// Upgrade: websocket
// Connaction: Upgrade
// Sec-Websocket-Accept: gfkjkfklgkglfuldfu

package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/sacOO7/gowebsocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New("ws://echo.websocket.org/")

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Printf("Connection OK")
	}

	socket.OnConnectError = func(err error, socker gowebsocket.Socket) {
		log.Printf("Error to connection: %v", err)
	}

	socket.OnPingReceived = func(_ string, socket gowebsocket.Socket) {
		log.Printf("Ping has resieved")
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Printf("Message: %v", message)
	}

	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Print("Connectionn closed")
			socket.Close()
			return
		case <-time.After(time.Second * 10):
			socket.SendText("Hello world")
		}
	}
}
