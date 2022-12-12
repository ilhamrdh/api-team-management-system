package config

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDatabase() *gorm.DB {
	value := url.Values{}
	value.Add("parseTime", "True")
	value.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%v)/%s?%s`, Config.Database.DBUser, Config.Database.DBPass, Config.Database.DBHost, Config.Database.DBPort, Config.Database.DBName, value.Encode())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Cannot connected database ", err)
	}
	return db
}
