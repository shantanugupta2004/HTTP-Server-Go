package tests

import (
	"testing"
	"http-server-go/database"
)

func TestBasicMath(t *testing.T) {
	if 2+2 != 4 {
		t.Error("Basic math failed")
	}
}

func TestDBConnection(t *testing.T) {
	database.Connect()

	if database.DB == nil {
		t.Fatal("Database connection failed: DB is nil")
	}
}