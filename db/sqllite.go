package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("sqlite", "sqlite/appdata.db")
		if err != nil {
			log.Fatalf("failed to open database: %v", err)
		}
	})
	return db
}

func Initialize() {
	_ = GetDB()
	m, err := migrate.New(
		"file://db/migrations",
		"sqlite://sqlite/appdata.db",
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("Database created or opened and migrations applied successfully!")

}
