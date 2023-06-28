package handler

import (
	"gozero-demo/template/demo2/internal/logic"
	"gozero-demo/template/demo2/internal/svc"
	"gozero-demo/template/demo2/internal/types"
	"gozero-demo/template/response" // ①
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Demo2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDemo2Logic(r.Context(), svcCtx)
		resp, err := l.Demo2(&req)
		response.Response(w, resp, err) //②

	}
}
