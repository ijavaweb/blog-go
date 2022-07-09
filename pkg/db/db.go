package db

import (
	"blog-go/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)
var DB *gorm.DB
func InitMySQL() {
	dsn := "root:123@tcp(localhost:3306)/blog"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logger.ErrorLogger.Error(err.Error())
		panic(err.Error())
	}
	sqlDB := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
	// Remove if not necessary after debug.
	DB.LogMode(true)
	// gorm v1 compatible settings.
	DB.Callback().Update().Remove("gorm:update_time_stamp")
	logger.InfoLogger.WithField("DB",db).Info("init db success")
}
