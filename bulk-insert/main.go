package main

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)
	/*
		blk, err := sqlx.NewBulkInserter(conn, "insert into user (id, name) values (?,?)")
		if err != nil {
			panic(err)
		}

		defer blk.Flush()

		blk.Insert(12, "test1")
		blk.Insert(22, "test2")

		blk.SetResultHandler(func(r sql.Result, err error) {
			if err != nil {
				logx.Error(err)
				return
			}
			fmt.Println(r.RowsAffected())
		})
	*/

	err := conn.TransactCtx(context.Background(), func(ctx context.Context, s sqlx.Session) error {
		r, err := s.ExecCtx(ctx, "insert into user (id, name) values (?, ?)", 5, "test1")
		if err != nil {
			return err
		}
		fmt.Println(r.RowsAffected())

		r, err = s.ExecCtx(ctx, "insert into user (id, name) values (?, ?)", 4, "test2")
		if err != nil {
			return err
		}
		fmt.Println(r.RowsAffected())
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
