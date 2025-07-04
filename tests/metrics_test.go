package tests

import (
	"net/http"
	"testing"
)

func TestMetricsEndpoint(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/metrics")
	if err != nil {
		t.Fatalf("Failed to GET /metrics: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	}
}
