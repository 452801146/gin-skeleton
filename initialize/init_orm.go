package initialize

import (
	"fmt"
	"gin_skeleton/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitGorm() {

	dbHost := g.Config.GetString("mysql.dbHost")
	dbPort := g.Config.GetString("mysql.dbPort")
	dbName := g.Config.GetString("mysql.dbName")
	dbUser := g.Config.GetString("mysql.dbUser")
	dbPass := g.Config.GetString("mysql.dbPass")

	newLogger := initGormLogger()

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True`, dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,      //禁用事务
		Logger:                 newLogger, //新日志
	})
	if err != nil {
		panic(err)
		return
	}
	baseDb, err := db.DB()
	if err != nil {
		panic(err)
		return
	}
	baseDb.SetMaxOpenConns(g.Config.GetInt("mysql.maxOpenConns"))
	baseDb.SetMaxIdleConns(g.Config.GetInt("mysql.maxIdleConns"))
	baseDb.SetConnMaxLifetime(time.Second * (g.Config.GetDuration("mysql.connMaxLifetime")))

	g.Orm = db
}

func initGormLogger() logger.Interface {

	// 日志写入文件
	writer := GetLogWriter("server.sqlLog")
	// 日志级别
	logModel := g.Config.GetInt("gorm.logLevel")

	newLogger := logger.New(
		log.New(writer, "\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,               // 慢 SQL 阈值
			LogLevel:      logger.LogLevel(logModel), // Log level
			Colorful:      false,                     // 是否开启彩色打印
		},
	)

	return newLogger

}
