package main

import (
	"fmt"
	"go/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
)

func main() {
	srv := rest.MustNewServer(rest.RestConf{
		Port: 8080,
	})

	srv.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/hello",
		Handler: handle,
	})

	defer srv.Stop()
	//httpx.SetErrorHandler仅在调用了httpx.Error处理响应时才有效
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *errors.CodeMsg:
			return http.StatusOK, xhttp.BaseResponse[types.Nil]{
				Code: e.Code,
				Msg:  e.Msg,
			}
		default:
			return http.StatusInternalServerError, nil

		}
	})
	srv.Start()
}

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloRespone struct {
	Msg string `json:"msg"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	var req HelloRequest
	if err := httpx.Parse(r, &req); err != nil {
		fmt.Println(111)
		httpx.Error(w, err)
		return
	}
	fmt.Println(req)

	if req.Name == "error" {
		//模拟参数错误
		fmt.Println(222)
		httpx.Error(w, errors.New(400, "参数错误"))
		return
	}

	fmt.Println(333)
	httpx.OkJson(w, HelloRespone{
		Msg: "hello " + req.Name,
	})
}
