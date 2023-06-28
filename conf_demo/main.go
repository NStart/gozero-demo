package main

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/zrpc"
)

type config struct {
	Etcd     discov.EtcdConf
	UserRpc  zrpc.RpcClientConf
	PortRpc  RpcClientConf
	OtherRpc RpcClientConf
}

type RpcClientConf struct {
	//要继承关键是定义`json:",optional,inherit"`
	Etcd discov.EtcdConf `json:",optional,inherit"`
}

func main() {
	var cfg config
	conf.LoadConfig("config.yaml", &cfg)

	fmt.Println(cfg.UserRpc.Etcd.Hosts)
	fmt.Println(cfg.PortRpc.Etcd.Hosts)

	var c logc.LogConf
	logc.MustSetup(c)

	logc.Info(context.Background(), "log")
}
