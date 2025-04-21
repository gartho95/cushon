package database

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Connect() (*gorm.DB, error) {
	var err error
	driver := os.Getenv("DB_DRIVER")

	slog.Info(fmt.Sprintf("Launching database connection for %s driver", driver))

	once.Do(func() {

		switch driver {
		case "sqlite":
			db, err = gorm.Open(sqlite.Open("cushon.db"), &gorm.Config{})
		case "mysql":
			dsn := os.Getenv("MYSQL_DSN")
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		case "postgres":
			dsn := os.Getenv("POSTGRES_DSN")
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		default:
			err = fmt.Errorf("unsupported DB_DRIVER: %s", driver)
		}
	})

	return db, err
}
