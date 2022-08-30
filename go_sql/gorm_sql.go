package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "fmt"
)

const (
    USERNAME = "demo"
    PASSWORD = "demo123"
    NETWORK = "tcp"
    SERVER = "127.0.0.1"
    PORT = 3306
    DATABASE = "demo"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}
}
