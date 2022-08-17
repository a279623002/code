package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

const (
	ADDRESS   = "tcp://114.116.44.156:1883"
)

var msg string

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

type MqttClient struct {
	c mqtt.Client
}

func newClient(name string) *MqttClient {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(ADDRESS).SetClientID(name)

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	return &MqttClient{c:c}
}

func (m *MqttClient) Pub(msg string) {
	token := m.c.Publish("testtopic/1", 0, false, msg)
	token.Wait()
}

func (m *MqttClient) Unsubscribe() {
	// 取消订阅
	if token := m.c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 断开连接
	m.c.Disconnect(250)
	time.Sleep(1 * time.Second)
}

// 从命令行获取客户端命名
func main() {
	if len(os.Args) < 2 {
		panic("need client")
	}
	client := os.Args[1]
	c := newClient(client)
	defer c.Unsubscribe()

	for {
		fmt.Print("send msg: ")
		fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		fmt.Println("you send msg:", msg)
		c.Pub(msg)
	}
	fmt.Println("done")
}
