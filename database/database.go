package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "root:root@tcp(localhost:8889)/workshop?parseTime=true"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)

	if err != nil {
		panic(err)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	Database = DbInstance{Db: db}
}
