package models

import (
	"fmt"
	"log"
	"page-ss/src/config"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// db连接
var db *gorm.DB

// Setup 初始化连接
func setupPool() {
	// db = newConnection()
	var dbURI string
	var dialector gorm.Dialector

	if config.Global.DbConfig.DbType == "mysql" {
		// logger.Log.Info("select mysql")
		dbURI = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
			config.Global.DbConfig.User,
			config.Global.DbConfig.Password,
			config.Global.DbConfig.Host,
			config.Global.DbConfig.Port,
			config.Global.DbConfig.DbName)
		fmt.Printf("dbURI=%v",dbURI)
		dialector = mysql.New(mysql.Config{
			DSN:                       dbURI, // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		})
	} else if config.Global.DbConfig.DbType == "postgres" {
		// logger.Log.Info("select postgres")
		dbURI = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
			config.Global.DbConfig.Host,
			config.Global.DbConfig.Port,
			config.Global.DbConfig.User,
			config.Global.DbConfig.DbName,
			config.Global.DbConfig.Password)
		dialector = postgres.New(postgres.Config{
			DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		})
	} else { // sqlite3
		// logger.Log.Info("select sqlite3")
		dbURI = fmt.Sprintf("test.db")
		dialector = sqlite.Open("test.db")
	}
	// Silent, Error, Warn, Info
	conn, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Print(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		// logger.Log.Error("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(10)                   // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(100)                  // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Second * 600) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

	//logger.Log.Info("setup new db connection......")
	db = conn
}

// GetDB 开放给外部获得db连接
func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		// logger.Log.Error("connect db server failed.")
		setupPool()
	}
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		setupPool()
	}

	return db
}


// Setup 初始化数据模型
func Setup() {
	setupPool()
	db := GetDB()
	// 自动迁移模式
	//db.AutoMigrate(&ParkingLot{}, &ParkingSpace{}, &Product{})
	db.AutoMigrate()
}
