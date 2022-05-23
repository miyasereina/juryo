package statements

import (
	"fmt"

	"juryo/dao"
)

type Image struct {
	Id        string `json:"id"`
	Md5       []byte `json:"md_5"`
	Fid       int64  `json:"fid"`
	Size      int32  `json:"size"`
	Width     int32  `json:"width"`
	Height    int32  `json:"height"`
	ImageType int32  `json:"image_type"`
	GroupId   int64  `json:"group_id"`
}

func ImageInit() {
	db := dao.MysqlConn()
	if !db.HasTable(&Image{}) {
		if err := db.CreateTable(&Image{}).Error; err != nil {
			panic(err)
		}
		fmt.Println("Table Image has been created")
	} else {
		db.AutoMigrate(&Image{})
		fmt.Println("Table Image existed")
	}
	dao.TimeSetting("image")
}
