package mysql

import (
	"antd-admin-golang/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error

	config := config.New()
	user := config.Mysql.User
	password := config.Mysql.Password
	host := config.Mysql.Host
	port := config.Mysql.Port
	database := config.Mysql.Database

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	db, err = gorm.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

}
