package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/api/handler"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/util/helper"
	cors "github.com/micro/micro/v3/service/api/server/http"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	mux := http.NewServeMux()
	// Initialize Server
	s := service.New(service.Name("user-srv"))

	mux.Handle("/", NewRPCHandler(s.Client()))

	log.Println("listen on: 9099")
	http.ListenAndServe(":9099", mux)
}

type rpcHandler struct {
	client   client.Client
}

type rpcRequest struct {
	Service  string
	Endpoint string
	Method   string
	Address  string
	Request  interface{}
}

func NewRPCHandler(c client.Client) handler.Handler {
	return &rpcHandler{c}
}

func (h *rpcHandler) String() string {
	return "internal/rpc"
}

func (h *rpcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		cors.SetHeaders(w, r)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	badRequest := func(description string) {
		e := errors.BadRequest("go.micro.rpc", description)
		w.WriteHeader(400)
		w.Write([]byte(e.Error()))
	}

	var service, endpoint, address string
	var request interface{}

	// response content type
	w.Header().Set("Content-Type", "application/json")

	ct := r.Header.Get("Content-Type")

	// Strip charset from Content-Type (like `application/json; charset=UTF-8`)
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}

	switch ct {
	case "application/json":
		var rpcReq rpcRequest

		d := json.NewDecoder(r.Body)
		d.UseNumber()

		if err := d.Decode(&rpcReq); err != nil {
			badRequest(err.Error())
			return
		}

		service = rpcReq.Service
		endpoint = rpcReq.Endpoint
		address = rpcReq.Address
		request = rpcReq.Request
		if len(endpoint) == 0 {
			endpoint = rpcReq.Method
		}

		// JSON as string
		if req, ok := rpcReq.Request.(string); ok {
			d := json.NewDecoder(strings.NewReader(req))
			d.UseNumber()

			if err := d.Decode(&request); err != nil {
				badRequest("error decoding request string: " + err.Error())
				return
			}
		}
	default:
		r.ParseForm()
		service = r.Form.Get("service")
		endpoint = r.Form.Get("endpoint")
		address = r.Form.Get("address")
		if len(endpoint) == 0 {
			endpoint = r.Form.Get("method")
		}

		d := json.NewDecoder(strings.NewReader(r.Form.Get("request")))
		d.UseNumber()

		if err := d.Decode(&request); err != nil {
			badRequest("error decoding request string: " + err.Error())
			return
		}
	}

	if len(service) == 0 {
		badRequest("invalid service")
		return
	}

	if len(endpoint) == 0 {
		badRequest("invalid endpoint")
		return
	}

	// create request/response
	var response json.RawMessage
	var err error
	req := client.NewRequest(service, endpoint, request, client.WithContentType("application/json"))

	// create context
	ctx := helper.RequestToContext(r)

	var opts []client.CallOption

	timeout, _ := strconv.Atoi(r.Header.Get("Timeout"))
	// set timeout
	if timeout > 0 {
		opts = append(opts, client.WithRequestTimeout(time.Duration(timeout)*time.Second))
	}

	// remote call
	if len(address) > 0 {
		opts = append(opts, client.WithAddress(address))
	}

	// remote call
	err = h.client.Call(ctx, req, &response, opts...)
	fmt.Println("request", req.Service(), err)
	if err != nil {
		ce := errors.Parse(err.Error())
		switch ce.Code {
		case 0:
			// assuming it's totally screwed
			ce.Code = 500
			ce.Id = "go.micro.rpc"
			ce.Status = http.StatusText(500)
			ce.Detail = "error during request: " + ce.Detail
			w.WriteHeader(500)
		default:
			w.WriteHeader(int(ce.Code))
		}
		w.Write([]byte(ce.Error()))
		return
	}

	b, _ := response.MarshalJSON()
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)
}

//
//func handlerRPC(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path == "/" {
//		w.Write([]byte("ok,this is the server ..."))
//		return
//	}
//	if r.URL.RequestURI() == "/favicon.ico" {
//		return
//	}
//	// 跨域处理
//	if origin := r.Header.Get("Origin"); cors[origin] {
//		w.Header().Set("Access-Control-Allow-Origin", origin)
//	} else if len(origin) > 0 && cors["*"] {
//		w.Header().Set("Access-Control-Allow-Origin", origin)
//	}
//
//	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-Token, X-Client")
//	w.Header().Set("Access-Control-Allow-Credentials", "true")
//
//	if r.Method == "OPTIONS" {
//		log.Println("method is options")
//		return
//	}
//
//	//if r.Method != "POST" {
//	//	log.Println("method not allowed")
//	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//	//	return
//	//}
//	handleJSONRPC(w, r)
//	return
//}
//
//func handleJSONRPC(w http.ResponseWriter, r *http.Request) {
//	service1, method := apid.PathToReceiver("go.micro.", r.URL.Path)
//	fmt.Println(service1, method)
//	br, _ := ioutil.ReadAll(r.Body)
//
//	request := json.RawMessage(br)
//
//	var response json.RawMessage
//	s := service.New()
//	req := s.Client().NewRequest("user-srv", "UserSrv.selectUser", &request)
//	ctx := apid.RequestToContext(r)
//	err := s.Client().Call(ctx, req, &response)
//	// make the call
//	if err != nil {
//		ce := errors.Parse(err.Error())
//		switch ce.Code {
//		case 0:
//			// assuming it's totally screwed
//			ce.Code = 500
//			ce.Id = service1
//			ce.Status = http.StatusText(500)
//			// ce.Detail = "error during request: " + ce.Detail
//			w.WriteHeader(500)
//		default:
//			w.WriteHeader(int(ce.Code))
//		}
//		w.Write([]byte(ce.Error()))
//		return
//	}
//	b, _ := response.MarshalJSON()
//	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
//	w.Write(b)
//}
