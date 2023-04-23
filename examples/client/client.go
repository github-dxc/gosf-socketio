package main

import (
	"log"
	"runtime"
	"time"

	gosocketio "github.com/github-dxc/gosf-socketio"
	"github.com/github-dxc/gosf-socketio/transport"
)

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c, err := gosocketio.Dial(
		gosocketio.GetUrl("localhost", 3811, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("/message", func(h *gosocketio.Channel, args Message) string {
		log.Println("Client Receive message: ", args)
		return "ok"
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
	//ack, err := c.Ack("/message", "hello", 5*time.Second)
	//log.Println(ack)
	//err = c.Emit("/message", "hello")
	//if err != nil {
	//	return
	//}
	time.Sleep(10 * time.Second)

	c.Close()

	select {}
}
