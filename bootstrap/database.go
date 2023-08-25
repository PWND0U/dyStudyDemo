package bootstrap

import (
	"dyStudyDemo/app/config"
	"dyStudyDemo/app/database"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

func getMysqlDsn() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config.C().Database.Mysql.Username,
		config.C().Database.Mysql.Password,
		config.C().Database.Mysql.Host,
		config.C().Database.Mysql.Port,
		config.C().Database.Mysql.DbName,
		config.C().Database.Mysql.Charset)
}

func SetupDatabase() {
	var dbConfig gorm.Dialector
	switch config.C().Database.Connection {
	case "mysql":
		dbConfig = mysql.New(mysql.Config{DSN: getMysqlDsn()})
	case "sqlite":
		dbConfig = sqlite.Open(config.C().Database.Sqlite.DbName)
	}
	database.Connect(dbConfig, gormLogger.Default.LogMode(gormLogger.Info))
	database.SQLDB.SetMaxOpenConns(config.C().Database.MaxOpenConnections)
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.C().Database.MaxIdleConnections)
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.C().Database.MaxLifeSeconds) * time.Second)
}
