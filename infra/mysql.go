package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectMysqlDB : connect to mysql database
func ConnectMysqlDB() (interface{}, error) {
	config := NewConfig()
	db, err := gorm.Open("mysql", config.Mysql.User+":"+config.Mysql.Password+"@"+config.Mysql.TCP+"/"+config.Mysql.DBName)
	return db, err
}
