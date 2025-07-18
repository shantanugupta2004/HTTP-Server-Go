package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"http-server-go/database"
	"http-server-go/models"
	"http-server-go/routes"
	"testing"
)

func init() {
	// Connect to the database and migrate tables for tests
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
}

func TestRegisterUser(t *testing.T) {
	router := routes.SetupRoutes()

	payload := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK && resp.Code != http.StatusCreated {
		t.Errorf("Expected 200 or 201, got %d", resp.Code)
	}
}
