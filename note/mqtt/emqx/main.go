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
	ClientID   = "emqx_test_client"
	USER_NAME = "admin"
	PASSWORD  = "public"

	Qos0 = 0 // 至多一次
	Qos1 = 1 // 至少一次
	Qos2 = 2 // 确保只有一次
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(ADDRESS).SetClientID(ClientID)

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

	//发布消息
	// qos是服务质量: ==1: 一次, >=1: 至少一次, <=1:最多一次
	// retained: 表示mqtt服务器要保留这次推送的信息，如果有新的订阅者出现，就会把这消息推送给它（持久化推送）
	token := c.Publish("testtopic/1", 0, false, "Hello World")
	token.Wait()

	time.Sleep(6 * time.Second)

	// 取消订阅
	if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 断开连接
	c.Disconnect(250)
	time.Sleep(1 * time.Second)
}