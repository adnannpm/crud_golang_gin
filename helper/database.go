package helper

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error){
	dialect := mysql.Open("root:root@tcp(localhost:3306)/supabank?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}