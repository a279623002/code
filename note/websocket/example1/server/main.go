package main

import (
	"container/list"
	"fmt"
	"net/http"
)

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	// 等待消息链表
	waitingList := list.New()
	// 如果获取事件为空，即等待消息，此处造成阻塞关键点
	// Wait for new message(s).
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch

	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("lastReceived"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/fetch", fetchHandler)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}