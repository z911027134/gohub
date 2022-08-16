package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger logger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
