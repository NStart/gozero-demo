package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

/*
type Request struct {
	Node   string `path:"node"`
	ID     int    `form:"id"`
	Header string `header:"X-Header"`
}
*/

type Request struct {
	Node   string            `path:"node"`
	Body   map[string]string `json:"body"`
	Header string            `header:"X-Header"`
}

var port = flag.Int("port", 3333, "the port to listen")

func handle(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	fmt.Println(111, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(2222)
	httpx.OkJson(w, req)
}

func main() {
	flag.Parse()

	svr := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port: *port,
	})
	defer svr.Stop()

	svr.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/nodes/:node",
		Handler: handle,
	})
	svr.Start()
}
