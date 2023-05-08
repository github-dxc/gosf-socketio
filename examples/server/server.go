package main

import (
	"fmt"
	gosocketio "github.com/github-dxc/gosf-socketio"
	"github.com/github-dxc/gosf-socketio/transport"
	"log"
	"net/http"
)

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		//go func() {
		//	time.Sleep(1 * time.Second)
		//	c.Emit("/message", Message{Text: "asdasdasd"})
		//c.Ack("/message", Message{Text: "asdasdasd"}, 5*time.Second)
		//}()
		log.Println("Connected")
		c.Join(c.Id())
	})
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		c.Leave(c.Id())
		log.Println("Disconnected")
	})

	server.On("/message", func(c *gosocketio.Channel, msg string) string {
		fmt.Println("Server Receive:", msg)
		return msg
	})

	serveMux := http.NewServeMux()
	namespace := "hello"
	serveMux.Handle("/socket.io/"+namespace, server)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(":3811", serveMux))
}
