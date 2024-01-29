package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DbInfo() string {
	env := GetEnv()
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_DATABASE, "utf8")

	return dsn
}

func newDb() *gorm.DB {
	dsn := DbInfo()

	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		db = newDb()
	}
	return db
}
