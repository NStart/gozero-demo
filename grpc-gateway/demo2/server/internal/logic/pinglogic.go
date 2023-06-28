package logic

import (
	"context"

	"gozero-demo/grpc-gateway/demo2/server/hello"
	"gozero-demo/grpc-gateway/demo2/server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *hello.Request) (*hello.Response, error) {
	// todo: add your logic here and delete this line

	return &hello.Response{
		Msg: "Pong",
	}, nil
}
