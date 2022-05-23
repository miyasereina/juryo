package statements

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"juryo/dao"
	"time"
)

type Month struct {
	gorm.Model
	Today time.Time `json:"today"`
	D     int       `gorm:"default: 0" json:"d"`
}

func MonthInit() {
	db := dao.MysqlConn()
	if !db.HasTable(&Month{}) {
		if err := db.CreateTable(&Month{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Month has been created")
	} else {
		db.AutoMigrate(&Month{})
		fmt.Println("Table Month existed")
	}
	dao.TimeSetting("month")
}
