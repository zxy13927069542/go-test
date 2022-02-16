package main

import (
	"bufio"
	"fmt"
	log "github.com/pion/ion-log"
	"net"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
	"time"
)

var (
	isNeedReconnect = true
	sessId          string
	isEstablish     = false
)

//
// Client
//  @Description: 与p90消息总线建立tcp长连接
//
type Client struct {
	conn           net.Conn
	onMessage      func(data []byte)
	onConnected    func()
	onDisconnected func(err error)
	establish      func() error
}

//
// NewClient
//  @Description: 新建一个tcp连接
//  @param addr 目标主机的socket ip:port
//
func NewClient(addr string) (*Client, error) {
	log.Infof("NewClient(), Connecting to %s...\n", addr)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Errorf("Failed to connect to server[%v]", addr)
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

//
// OnMessage
//  @Description: 绑定消息处理函数
//  @receiver c
//
func (c *Client) OnMessage(handler func(data []byte)) {
	c.onMessage = handler
}

//
// OnConnected
//  @Description: 绑定心跳处理函数
//
func (c *Client) OnConnected(handler func()) {
	c.onConnected = handler
}

//
// OnDisconnected
//  @Description: 绑定连接重连函数
func (c *Client) OnDisconnected(handler func(err error)) {
	c.onDisconnected = handler
}

func (c *Client) OnEstablish(handler func() error) {
	c.establish = handler
}

//
// listen
//  @Description: 接收tcp数据和发送心跳
//  @receiver c
//
func (c *Client) listen() {
	//reader := bufio.NewReader(c.conn)
	//buffer := make([]byte, 1024 * 10)
	//var packet []byte

	err := c.establish()
	if err != nil {
		c.conn.Close()
		c.onDisconnected(err)
		return
	}

	//  开始心跳
	go c.onConnected()

	reader := bufio.NewReader(c.conn)
	buffer := make([]byte, 1024*10)
	for {
		var packet []byte
		n, err := reader.Read(buffer)
		if err != nil {
			c.conn.Close()
			log.Infof("ReadBytes err: %v", err)
			c.onDisconnected(err)
			return
		}

		log.Infof("read bytes....., n=%v", n)

		if n > 0 {
			packet = append(packet, buffer[:n]...)
			//  判断报文是否结束
			if isMatch, _ := regexp.Match(`\{.*}`, packet); isMatch {
				//  截取数据
				compile, _ := regexp.Compile(`\{.*}`)
				data := compile.Find(packet)

				log.Infof("receive data: %s", string(data))

				go c.onMessage(data)
			}
		}

	}
}

// Send
//  @Description: 通过tcp连接发送消息
//  @param message 待发送的消息
//
func (c *Client) Send(message string) error {
	log.Infof("Send: message=%v", message)
	return c.SendBytes([]byte(message))
}

// Send bytes to client
func (c *Client) SendBytes(b []byte) error {
	// log.Infof("SendBytes: data=[% 02x]", b)

	_, err := c.conn.Write(b)
	if err != nil {
		c.conn.Close()
		log.Infof("c.conn.Write err: %v", err)
	}
	return err
}

func start() {
	host := "10.7.3.197"
	port := "30002"
	addr := host + ":" + port
	client, err := NewClient(addr)
	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
	} else {
		isNeedReconnect = false

		//  注册消息处理函数
		client.OnMessage(func(data []byte) {
			log.Infof("%s", string(data))
		})

		client.OnEstablish(func() error {
			//  开始发送心跳
			err, str := establish(client)
			if err != nil {
				return err
			}

			if len(str) > 0 {
				log.Infof("connection established! sessId: %s", str)
				sessId = str
			}
			return nil
		})

		//  注册心跳发送函数
		client.OnConnected(func() {
			//  开始发送心跳
			startHeartBeat(client)
		})

		//  注册重连函数
		client.OnDisconnected(func(err error) {
			isNeedReconnect = true
			log.Infof("Client disconnected. isNeedReconnect=%v, err=%v", isNeedReconnect, err)
		})
	}

	go client.listen()
}

//
// establish
//  @Description: 往tcp连接发送/establish的http报文
//
func establish(c *Client) (error, string) {
	//  构造http报文
	req, err := http.NewRequest(http.MethodPost, "http://10.7.3.197:30002/establish", nil)
	req.Header.Add("Host", "localhost")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Content-Length", "0")
	if err != nil {
		log.Errorf("Create /establish request failed! reason: %v", err)
		return err, ""
	}

	//  将http报文转成字节流
	byteReq, err := httputil.DumpRequest(req, false)
	if err != nil {
		log.Errorf("DumpRequest() Failed! Reason: %v", err)
		return err, ""
	}

	//log.Infof("%s", byteReq)
	err = c.SendBytes(byteReq)
	if err != nil {
		return err, ""
	}

	reader := bufio.NewReader(c.conn)
	buffer := make([]byte, 1024*10)
	var packet []byte
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			log.Infof("ReadBytes err: %v", err)
			return err, ""
		}

		if n > 0 {
			packet = append(packet, buffer[:n]...)
			//  判断报文是否结束
			if isMatch, _ := regexp.Match(`\{.*}`, packet); isMatch {
				return nil, getSessId(packet)
			}
		}
	}
}

func getSessId(packet []byte) string {
	message := string(packet)
	index := strings.Index(message, "key")
	startIndex := index + len(`key":"`)
	endIndex := len(message) - len(`"}`)
	return message[startIndex:endIndex]
}

//
// heartbeat
//  @Description: 往tcp连接发送/heartbeat的http报文
//
func heartbeat(c *Client) {
	//  构造http报文
	req, err := http.NewRequest(http.MethodPost, "http://10.7.3.197:30002/heartbeat", nil)
	req.Header.Add("Host", "localhost")
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Content-Length", "0")
	if err != nil {
		log.Errorf("Create /heartbeat request failed! reason: %v", err)
		return
	}

	//  将http报文转成字节流
	byteReq, err := httputil.DumpRequest(req, false)
	if err != nil {
		log.Errorf("DumpRequest() Failed! Reason: %v", err)
		return
	}

	//log.Infof("%s", byteReq)
	err = c.SendBytes(byteReq)
	if err != nil {
		return
	}

	return
}

//
// startHeartBeat
//  @Description: 定时发送心跳
//
func startHeartBeat(c *Client) {
	for {
		heartbeat(c)
		time.Sleep(time.Second * 30)
	}
}

func main() {
	log.Infof("startToProxyMode")

	isNeedReconnect = true
	for {
		if isNeedReconnect {
			log.Infof("\n--------------connecting to server---------------")
			log.Infof("startToProxyMode isNeedReconnect=%v", isNeedReconnect)

			start()
		}

		time.Sleep(time.Second * 5)
	}
}
