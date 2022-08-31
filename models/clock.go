package models

import (
	"context"
	"juryo/dao"
	"juryo/models/statements"
	"strconv"
	"time"
)

var ctx = context.Background()

func Insert(user string, dur time.Duration) (string, error) {
	db := dao.MysqlConn()
	clock := statements.Clock{}
	db.Table("clock").Last(&clock)
	if clock.Name == user && clock.CreatedAt.Day() == time.Now().Day() {
		return "笨蛋" + user + "今天已经打过卡了哦", nil
	}
	err := db.Table("clock").Create(&statements.Clock{
		Name: user,
	}).Error
	rediscli := dao.RedisClient
	timeChannel := time.After(dur)
	auto := rediscli.Get(ctx, "auto").Val()
	e := rediscli.Get(ctx, "auto").Err()
	if e != nil {
		return "", e
	}
	if auto == "" {
		return "1", err
	}
	v, _ := strconv.Atoi(auto)
	go func() {
		select {
		case <-timeChannel:
			aut := rediscli.Get(ctx, "auto").Val()
			r := rediscli.Get(ctx, "auto").Err()
			if r != nil {
				panic(r)
			}
			if aut == "" {
				er := rediscli.SetEX(ctx, "auto", "1", time.Hour*24).Err()
				if er != nil {
					panic(er)
				}
			} else {
				v, _ := strconv.Atoi(auto)
				er := rediscli.SetEX(ctx, "auto", v+1, time.Hour*24).Err()
				if er != nil {
					panic(er)
				}

			}
		}
	}()
	str := strconv.Itoa(v + 1)
	return user + "打卡第" + str + "天", err
}
