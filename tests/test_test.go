package tests

import (
	sqlcommentergorm "github.com/a631807682/sqlcommenter-gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	dbDSN := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Logger = db.Logger.LogMode(logger.Info)
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	// use plugin
	db.Use(sqlcommentergorm.Default())
	DB = db
}
