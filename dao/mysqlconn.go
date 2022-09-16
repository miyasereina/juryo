package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func MysqlInit() {
	dbName := "bot"
	user := "bot"
	password := "@Reina0526"
	port := "127.0.0.1"
	dbInfo := user + ":" + password + "@tcp(" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"

	//connect
	var err error
	DB, err = gorm.Open("mysql", dbInfo)
	DB.SingularTable(true)
	DB.DB().SetMaxOpenConns(50)
	DB.DB().SetMaxIdleConns(20)
	if err != nil {
		panic(err)
	}
}
func MysqlConn() *gorm.DB {
	return DB
}

func TimeSetting(tableName string) {
	db := MysqlConn()
	//插入时自动更新created_at字段
	db.Exec("ALTER TABLE " + tableName + " MODIFY created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL;")
	//更新时自动更新updated_at字段
	db.Exec("ALTER TABLE " + tableName + " MODIFY updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL;")

}
