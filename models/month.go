package models

import (
	"fmt"
	"juryo/dao"
	"time"
)

type month struct {
	ID    int       `json:"id"`
	Today time.Time `json:"today"`
	D     int       ` json:"d"`
}

func Flush() (re string) {
	m := &month{}
	this := &month{}
	d := make([]int, 1)
	aver := 0
	db := dao.MysqlConn().Table("month")
	db.Last(&m).Update()
	this.Today = time.Now()
	if m != nil {
		this.D = int(this.Today.Sub(m.Today).Hours() / 24)
		db.Raw("SELECT SUM(d) as d FROM month").Pluck("SUM(d) as d", &d)
		aver = (d[0] + this.D) / m.ID
		if this.D-aver >= 7 || aver-this.D >= 7 {
			re = "偏差值超过七天了捶捶妈妈，近期一定注意身体健康"
		}
		dur := int64(aver * 24)
		re = "预计下次为" + this.Today.Add(time.Duration(dur)*time.Hour).Format("2006-01-02") + re
		fmt.Println(this.D, this.Today.Sub(m.Today), this.Today.Format("2006-01-02"), m.Today.Format("2006-01-02"))
	}
	db.Create(&this)
	re = "本次为" + this.Today.Format("2006-01-02") + "," + re + "。辛苦妈妈了，抱抱妈妈(づ′▽`)づ"
	return re
}

func Add(ti []int) {
	db := dao.DB.Table("month")
	fmt.Println(ti)
	db.Create(&month{
		Today: time.Date(ti[0], time.Month(ti[1]), ti[2], 0, 0, 0, 0, time.UTC),
		D:     ti[3],
	})

}
