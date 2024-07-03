package database

import (
	"time"

	"gorm.io/gorm"
)

var DBManager *DatabaseManager

// DatabaseManager To Handle Multiple Database Connections
type DatabaseManager struct {
	SqliteDB *gorm.DB
}

// SqlLiteConfig represents configuration for sqlite database
type SqlLiteConfig struct {
	Reset      bool
	Tables     *[]string
	GCInterval *time.Duration

	Database        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// Initialize DBManager
func init() {
	DBManager = &DatabaseManager{}
}
