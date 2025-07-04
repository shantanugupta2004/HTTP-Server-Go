package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"http-server-go/routes"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := routes.SetupRoutes()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 but got %d", rr.Code)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, rr.Body.String())
	}
}
