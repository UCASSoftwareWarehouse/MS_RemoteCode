package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"remote_code/config"
)

var Db *gorm.DB

func InitGorm() {
	var err error
	conf := config.Conf
	uri := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.MysqlAddr, conf.Database)
	log.Printf("gorm uri=%+v", uri)
	Db, err = gorm.Open("mysql", uri)
	if err != nil {
		log.Fatalf("initGorm err:%+v", err)
	}
}

func GetDb() *gorm.DB {
	if Db == nil {
		InitGorm()
	}
	return Db
}
