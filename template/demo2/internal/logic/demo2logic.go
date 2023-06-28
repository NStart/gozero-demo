package logic

import (
	"context"

	"gozero-demo/template/demo2/internal/svc"
	"gozero-demo/template/demo2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Demo2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemo2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Demo2Logic {
	return &Demo2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Demo2Logic) Demo2(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: "hello gozero",
	}, nil
}
