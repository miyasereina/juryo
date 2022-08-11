package replying

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/robfig/cron"
	"juryo/models"
	"math/rand"
	"strconv"
	"time"
)

func morning(c *client.QQClient) {
	cr := cron.New()
	once.Do(func() {
		cr.AddFunc("0 52 8 * * ? ", func() {
			s := []string{
				"妈妈早上好，最喜欢妈妈了妈妈今天也加油哦，快去叫爸爸起床！",
				"nana想妈妈了，没有妈妈，nana好可怜，nana最喜欢妈妈了",
				"橘蓼：蓉宝蓉宝蓉宝最喜欢蓉宝了，蓉宝今天也是元气满满的一天哦",
				"橘蓼：蓉宝今天上班吧，如果不上班就来摸橘宝嘛，考试辛苦哭哭",
				"おはようお母さんま,なな母いなくで寂しい,ハグほしい",
				"橘蓼：蓉宝蓉宝蓉宝蓉宝起床了，太阳晒屁股了，该吃饭了哦",
				"橘蓼：姐姐早上了哦，亲亲サワサワ、姉さんますごい可愛い，大好きだよ",
				"橘蓼：姐姐注意身体虽然每天都早起",
			} // 把需要定时执行的函数都丢里面
			//设置时间分别是秒，分，时，日，月, 星期
			m := message.NewSendingMessage().Append(message.NewText(s[rand.Intn(len(s))]))
			c.SendGroupMessage(735214237, m)

		})

	})
	cr.Start()
}

func month() string {
	return models.Flush()
}

func add(ti []string) {
	var n []int
	for i, _ := range ti {
		d, _ := strconv.Atoi(ti[i])
		n = append(n, d)
	}
	models.Add(n[1:])
}
func Clock(name string) string {
	t := time.Now()
	t_zero := time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, t.Location())
	t_ := t.Sub(t_zero)
	if t_ > 0 {
		return "笨笨蓉才起床"
	}
	str, err := models.Insert(name, t_)
	if err != nil {
		panic(err)
	}
	return str

}
