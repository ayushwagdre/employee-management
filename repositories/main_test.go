package repository

import (
	"os"
	"practice/config"
	"practice/lib/db"
	"practice/test_helpers"
	"testing"
)

func TestMain(m *testing.M) {
	config := config.NewConfig()
	test_helpers.SetupDB(config)
	testResult := m.Run()
	test_helpers.ClearDataFromPostgres(db.Get().Get())
	os.Exit(testResult)
}
