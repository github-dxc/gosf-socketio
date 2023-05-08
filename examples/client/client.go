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
		gosocketio.GetNamespaceUrl("localhost", 3811, "hello", false),
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
	a := `{
  "code": 0,
  "data": {
    "nickname": "akaa",
    "id": 2511730347609102,
    "sec_uid": "MS4wLjABAAAA-4hQcqDZEsKZuoHg3vmI1s1Us-YMrrAdfMQaC-7VXXM_VbpaNbf-NSQJ-SR_nH7F",
    "city": "",
    "avatar_thumb": "https:\/\/p11.douyinpic.com\/aweme\/100x100\/aweme-avatar\/mosaic-legacy_3792_5112637127.jpeg?from=3067671334"
  },
  "msg": "获取成功"
}`
	ack, err := c.Ack("/message", a, 5*time.Second)
	log.Println(ack)
	//err = c.Emit("/message", "hello")
	//if err != nil {
	//	return
	//}
	time.Sleep(10 * time.Second)

	c.Close()

	select {}
}
