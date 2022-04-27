package main

import (
	"fmt"
	"net/http"
)

func print(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	fmt.Println(r.URL)
}

func main() {
	http.HandleFunc("/print", print)
	if err := http.ListenAndServe(":27963", nil); err != nil {
		fmt.Println(err)
		return
	}
}