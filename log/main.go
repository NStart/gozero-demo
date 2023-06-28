package main

import (
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	//输出到控制台
	//logc.Info(context.Background(), "hello world")

	//输出到日志文件
	/*
		var cfg logx.LogConf
		_ = conf.FillDefault(&cfg)
		cfg.Mode = "file"

		logc.MustSetup(cfg)
		defer logc.Close()
		logc.Info(context.Background(), "hello world")
	*/

	//带上额外信息
	//logc.Infow(context.Background(), "hello world", logc.Field("key", "value"))

	//所有日志带上默认的key和value
	/*
		ctx := logx.ContextWithFields(context.Background(), logx.Field("path", "user/info"))

		logc.Infow(ctx, "hello world")
		logc.Error(ctx, "error log")
	*/

	//可打印exec的位置
	exec()
}

func exec() error {
	logx.WithCallerSkip(1).Info("exec info")
	return nil
}
