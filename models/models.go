package models

import (
	"juryo/dao"
	"juryo/models/statements"
	"time"
)

func TableInit() {
	statements.ImageInit()
	statements.MonthInit()
	//add()
}

func add() {

	db := dao.MysqlConn()
	db.Table("month").Create(&month{
		Today: time.Date(2022, 1, 10, 0, 0, 0, 0, time.UTC),
		D:     0,
	}).Create(&month{
		Today: time.Date(2022, 2, 14, 0, 0, 0, 0, time.UTC),
		D:     35,
	}).Create(&month{Today: time.Date(2022, 3, 18, 32, 0, 0, 0, time.UTC),
		D: 32}).Create(&month{Today: time.Date(2022, 4, 15, 0, 0, 0, 0, time.UTC),
		D: 28})
}
