package gosocketio

import (
	"net"
	"strconv"
	"strings"

	"github.com/github-dxc/gosf-socketio/transport"
)

const (
	webSocketProtocol       = "ws://"
	webSocketSecureProtocol = "wss://"
	socketioUrl             = "/socket.io/?EIO=3&transport=websocket"
)

/*
*
Socket.io client representation
*/
type Client struct {
	methods
	Channel
}

/*
*
Get ws/wss url by host and port
*/
func GetUrl(host string, port int, secure bool) string {
	var prefix string
	if secure {
		prefix = webSocketSecureProtocol
	} else {
		prefix = webSocketProtocol
	}
	return prefix + net.JoinHostPort(host, strconv.Itoa(port)) + socketioUrl
}

// GetNamespaceUrl
//
//	@Description: 获取带命名空间的url
//	@param host
//	@param port
//	@param namespace 命名空间
//	@param secure
//	@return string
func GetNamespaceUrl(host string, port int, namespace string, secure bool) string {
	var prefix string
	if secure {
		prefix = webSocketSecureProtocol
	} else {
		prefix = webSocketProtocol
	}
	i := strings.Index(socketioUrl, "?")
	url := socketioUrl[:i] + namespace + socketioUrl[i:]
	return prefix + net.JoinHostPort(host, strconv.Itoa(port)) + url
}

/*
*
connect to host and initialise socket.io protocol

The correct ws protocol url example:
ws://myserver.com/socket.io/?EIO=3&transport=websocket

You can use GetUrlByHost for generating correct url
*/
func Dial(url string, tr transport.Transport) (*Client, error) {
	c := &Client{}
	c.initChannel()
	c.initMethods()

	var err error
	c.conn, err = tr.Connect(url)
	if err != nil {
		return nil, err
	}

	go inLoop(&c.Channel, &c.methods)
	go outLoop(&c.Channel, &c.methods)
	go pinger(&c.Channel)

	return c, nil
}

/*
*
Close client connection
*/
func (c *Client) Close() {
	closeChannel(&c.Channel, &c.methods)
}
