package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

//连接数据库

func Connect(dbConfig gorm.Dialector, _logger gormLogger.Interface) {
	var err error
	db, err := gorm.Open(dbConfig, &gorm.Config{Logger: _logger})
	if err != nil {
		fmt.Println(err.Error())
	}
	DB = db
	s, err := db.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
	SQLDB = s
}
