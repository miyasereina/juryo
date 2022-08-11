package statements

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"juryo/dao"
)

type Clock struct {
	gorm.Model
	Name string `json:"name"`
}

func ClockInit() {
	db := dao.MysqlConn()
	if !db.HasTable(&Clock{}) {
		if err := db.CreateTable(&Clock{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Month has been created")
	} else {
		db.AutoMigrate(&Clock{})
		fmt.Println("Table Clock existed")
	}
	dao.TimeSetting("clock")
}
