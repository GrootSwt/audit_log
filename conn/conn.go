package conn

import (
	"audit_log/common"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// 连接数据库
func GetDB() *gorm.DB {
	dbModel := common.ConfigInformation.DBConfigInfo
	host := dbModel.Host
	port := dbModel.Port
	user := dbModel.User
	password := dbModel.Password
	dbName := dbModel.DBName
	db, err := gorm.Open("mysql", user+":"+password+"@("+host+":"+port+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接出现错误：", err)
	}
	//	开启日志
	db.LogMode(true)
	return db
}
