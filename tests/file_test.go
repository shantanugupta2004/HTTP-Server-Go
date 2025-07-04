package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"http-server-go/database"
	"http-server-go/models"
	"http-server-go/routes"
)

var authToken string

func init() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.File{})

	router := routes.SetupRoutes()

	// Register a user
	registerPayload := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	body, _ := json.Marshal(registerPayload)

	req := httptest.NewRequest("POST", "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Login the user
	loginPayload := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	body, _ = json.Marshal(loginPayload)
	req = httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Parse token from response
	var resBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &resBody)
	authToken = resBody["token"].(string)
}

func TestUploadFile(t *testing.T) {
	router := routes.SetupRoutes()

	// Create fake file
	filePath := "./test.txt"
	os.WriteFile(filePath, []byte("This is test content"), 0644)
	defer os.Remove(filePath)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(filePath))
	file, _ := os.Open(filePath)
	defer file.Close()
	io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+authToken)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.Code)
	}
}

func TestListFiles(t *testing.T) {
	router := routes.SetupRoutes()

	req := httptest.NewRequest("GET", "/getFiles", nil)
	req.Header.Set("Authorization", "Bearer "+authToken)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.Code)
	}
}
