package main

import (
	"domain/apid"
	"encoding/json"
	"github.com/micro/go-micro/v2"
	microErrors "github.com/micro/go-micro/v2/errors"
	"log"
	"net/http"
	"strconv"
	pb "user-srv/proto"
)

var (
	cors = map[string]bool{"*": true}
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlerRPC)

	log.Println("listen on: 9099")
	http.ListenAndServe(":9099", mux)
}

func handlerRPC(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("ok,this is the server ..."))
		return
	}
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}
	// 跨域处理
	if origin := r.Header.Get("Origin"); cors[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else if len(origin) > 0 && cors["*"] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-Client")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		log.Println("method is options")
		return
	}

	if r.Method != "POST" {
		log.Println("method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	handleJSONRPC(w, r)
	return
}

func handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	srv, method := apid.PathToReceiver("go.micro.", r.URL.Path)

	service := micro.NewService()
	service.Init()
	c := service.Client()
	// Create new request to service go.micro.srv.example, method Example.Call
	var rr interface{}
	rr = &pb.SelectUserReq{
		Id: 1,
	}
	req := c.NewRequest(srv, method, rr)

	ctx := apid.RequestToContext(r)

	var rsp interface{}
	rsp = &pb.SelectUserResp{}

	// Call service
	if err := c.Call(ctx, req, rsp); err != nil {
		ce := microErrors.Parse(err.Error())
		switch ce.Code {
		case 0:
			// assuming it's totally screwed
			ce.Code = 500
			ce.Id = srv
			ce.Status = http.StatusText(500)
			// ce.Detail = "error during request: " + ce.Detail
			w.WriteHeader(500)
		default:
			w.WriteHeader(int(ce.Code))
		}
		w.Write([]byte(ce.Error()))
		return
	}

	b, _ := json.Marshal(rsp)
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)
}
