package database

import (
	"time"

	"github.com/Raman5837/kafka.go/base/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect To Database
func EstablishConnection() error {

	config := SqlLiteConfig{
		MaxIdleConns:    25,
		MaxOpenConns:    100,
		Reset:           false,
		Database:        "./.sqlite3",
		ConnMaxLifetime: 20 * time.Minute,
	}

	var connectionErr error
	gormConfig := &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	DBManager.SqliteDB, connectionErr = gorm.Open(sqlite.Open(config.Database), gormConfig)

	if connectionErr != nil {
		utils.Logger.Fatal(connectionErr, "Something Went Wrong While Connection To Sqlite3 DB.")
	}

	sqlite, _ := DBManager.SqliteDB.DB()
	sqlite.SetMaxIdleConns(config.MaxIdleConns)
	sqlite.SetMaxOpenConns(config.MaxOpenConns)
	sqlite.SetConnMaxLifetime(config.ConnMaxLifetime)

	if config.Reset {
		DBManager.ResetTable(*config.Tables)
	}

	return connectionErr

}
