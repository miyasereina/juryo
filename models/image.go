package models

import (
	"github.com/Mrs4s/MiraiGo/message"
	"juryo/dao"
	"juryo/models/statements"
)

func Upload(e *message.GroupImageElement, group_id int64) {
	db := dao.MysqlConn()
	eob := statements.Image{
		Id:        e.ImageId,
		Size:      e.Size,
		Width:     e.Width,
		Height:    e.Height,
		ImageType: e.ImageType,
		Md5:       e.Md5,
		Fid:       e.FileId,
		GroupId:   group_id,
	}
	db.Table("image").Create(&eob)

}
func Get(group_id int64) statements.Image {
	db := dao.MysqlConn()
	eob := statements.Image{}
	db.Table("image").Where("group_id=?", group_id).Order("rand()").First(&eob)
	return eob
}
