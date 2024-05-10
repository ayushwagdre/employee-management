package test_helpers

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"practice/config"
	"practice/lib/db"
	"strings"
	"time"

	"gorm.io/gorm"
)

var testDbClient db.DB

func SetupDB(config *config.Config) {
	testDbClient = db.NewDB(&db.DBOpts{
		URL:                   config.DB.URL,
		MaxIdleConnection:     10,
		MaxActiveConnection:   10,
		MaxConnectionLifetime: time.Hour,
		DriverName:            "pgx",
	})

	testDbClient.Connect()

	sqlDB, err := db.Get().Get().DB()
	if err != nil {
		panic(err)
	}
	loadDBSchema(sqlDB)

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
	truncateAllTables(db)
}

func truncateAllTables(db *gorm.DB) error {
	tableNames, err := loadTableNames(db)
	if err != nil {
		return err
	}

	if err := db.Exec("TRUNCATE TABLE " + strings.Join(tableNames, ", ")).Error; err != nil {
		return err
	}
	return nil
}

func loadTableNames(db *gorm.DB) ([]string, error) {
	rows, err := db.Raw("select table_name from information_schema.tables where table_schema='public'").Rows()
	if err != nil {
		return nil, err
	}

	var tableNames []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tableNames = append(tableNames, name)
	}

	return tableNames, nil
}
