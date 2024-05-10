package repository

import (
	"os"
	"practice/config"
	"practice/test_helpers"
	"testing"
)

func TestMain(m *testing.M) {
	config := config.NewConfig()
	test_helpers.SetupDB(config)
	testResult := m.Run()
	os.Exit(testResult)
}
