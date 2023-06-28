package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/rest/httpc"
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

var domain = flag.String("domain", "http://localhost:3333", "the domain to request")

func main() {
	flag.Parse()

	req := Request{
		Node: "foo",
		Body: map[string]string{
			"foo": "testfoo",
			"bar": "testbar",
		},
		Header: "foo-header",
	}

	resp, err := httpc.Do(context.Background(), http.MethodPost, *domain+"/nodes/:node", req)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, resp.Body)
}
