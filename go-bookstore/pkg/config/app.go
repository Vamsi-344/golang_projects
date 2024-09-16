package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open(), "admin:password/simplerest?charset=utf8")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
