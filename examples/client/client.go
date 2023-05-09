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
		gosocketio.GetUrl("localhost", 8828, false),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Fatal(err)
	}

	err = c.On("message", func(h *gosocketio.Channel, args Message) string {
		log.Println("Client Receive message: ", args)
		return "ok"
	})

	err = c.On("ReceiveMsg", func(h *gosocketio.Channel, args string) string {
		log.Println("Client ReceiveMsg: ", args)
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
	a := `{"namespace":"/","event":"ReceiveMsg","data":{"id":2511730347609102,"group_id":1}}`
	ack, err := c.Ack("SendMsg", a, 5*time.Second)
	log.Println(ack)
	//err = c.Emit("/message", "hello")
	//if err != nil {
	//	return
	//}
	time.Sleep(10 * time.Second)

	c.Close()

	select {}
}
