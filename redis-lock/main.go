package main

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	conf := redis.RedisConf{
		Host: "127.0.0.1:6379",
		Type: "node",
		Pass: "123456",
		Tls:  false,
	}
	rds := redis.MustNewRedis(conf)
	lock := redis.NewRedisLock(rds, "test")

	//设置过期时间
	lock.SetExpire(10)

	//尝试获取锁
	acquire, err := lock.Acquire()

	switch {
	case err != nil:
		//deal err
		fmt.Println(err)
	case acquire:
		defer lock.Release() //释放锁
		//业务逻辑
		fmt.Println(111)
	case !acquire:
		//没有拿到锁 wait?
		fmt.Println(222)
	}
}
