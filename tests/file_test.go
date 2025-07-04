package tests

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"http-server-go/routes"
	"testing"
)

func TestUploadFile(t *testing.T) {
	router := routes.SetupRoutes()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	filePath := "testdata/sample.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Unable to open test file: %v", err)
	}
	defer file.Close()

	part, _ := writer.CreateFormFile("file", filepath.Base(filePath))
	io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.Code)
	}
}

func TestListFiles(t *testing.T) {
	router := routes.SetupRoutes()
	req := httptest.NewRequest("GET", "/files", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.Code)
	}
}
