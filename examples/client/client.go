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
	a := `{"code":200,"msg":"success","data":{"id":11,"created_at":"2023-05-04T12:00:28.48+08:00","updated_at":"2023-05-04T12:00:28.48+08:00","account_id":3,"conversation_id":"0:1:5774988301:17972723560","conversation_short_id":7226277271248584971,"ticket":"UtcSKZa36lKiYUFN36XYatlZvhAMBqjaSe4DNWbftExKPqx3a0yxzAnzcDUpKKJtc3GexfBIuIzhLhc635R4UPwrXOWkwLnLY4FycOb5pNR9X4BVX0mj8HYhtEKdEMoefiWRUMcfvCJsrNecavf3NHT2naGEMwjaMN9Lv5CHa79Mssc4jb3guR9XiR5S6aMAkQ3KZ6DY8rMs","info_version":1682498800,"read_index":1682498805630707,"owner":5774988301,"friend_uid":17972723560,"friend_avatar_url":"https://sf6-cdn-tos.toutiaostatic.com/img/user-avatar/6c04365a22e14ea973e6bda759d09169~300x300.image","friend_screen_name":"莫天凝","conversation_ext":"[{\"key\":\"s:s_aid\",\"value\":\"13\"}]"}}`
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
