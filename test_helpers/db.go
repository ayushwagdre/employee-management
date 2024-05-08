package test_helpers

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"practice/config"
	"practice/lib/db"
	"time"

	"gorm.io/gorm"
)

var testDbClient db.DB

func SetupDB(config *config.Config) {
	var err error
	testDbClient = db.NewDB(&db.DBOpts{
		URL:                   config.DB.URL,
		MaxIdleConnection:     10,
		MaxActiveConnection:   10,
		MaxConnectionLifetime: time.Hour,
		DriverName:            "nrpgx",
	})

	testDbClient.Connect()
	if err != nil {
		panic(err)
	}
}

func loadDBSchema(sqlDB *sql.DB) {
	defaultSqlFilePath, _ := filepath.Abs("../db_migrations/db/structure.sql")
	fmt.Println(filepath.Join(filepath.Dir(defaultSqlFilePath), "structure.sql"))

	sqlContent, ioErr := os.ReadFile(filepath.Join(filepath.Dir(defaultSqlFilePath), "structure.sql"))
	if ioErr != nil {
		panic(ioErr)
	}

	_, err := sqlDB.Exec(string(sqlContent))
	if err != nil {
		panic(err)
	}
}

func ClearDataFromPostgres(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	loadDBSchema(sqlDB)
}
