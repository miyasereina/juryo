package cron

import (
	"fmt"
	"github.com/robfig/cron"
)

func Gron() *cron.Cron {

	c := cron.New()
	// 把需要定时执行的函数都丢里面
	// 设置时间分别是秒，分，时，日，月, 星期
	c.AddFunc("* * * * * *", func() {
		fmt.Println(1)
	})
	return c

}
