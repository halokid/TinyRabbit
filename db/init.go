package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

// fixme: 更改mysql配置
const (
  dbHost = "8.8.8.8"
  dbPort = "3306"
  dbUser = "root"
  dbPwd = "xxxx"
  dbName = "tinyrabbit"
)

func init() {
  var err error
  Db, err = gorm.Open("mysql", dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" +
                      dbName + "?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic(err)
  }
}






