package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeNow(t *testing.T) {
	// Test data
	testData := TimeNowInputs{
		Format: "RFC3339",
	}

	// Convert test data to JSON
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	// Create request
	req := httptest.NewRequest("POST", "/timeNow", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Execute the TimeNow function directly
	TimeNow(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check content type
	expectedContentType := "application/json"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected content type %s, got %s", expectedContentType, contentType)
	}

	// Parse response
	var responseData TimeNowOutputs
	if err := json.NewDecoder(w.Body).Decode(&responseData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Validate that the time is in RFC3339 format (current implementation)
	parsedTime, err := time.Parse(time.RFC3339, responseData.Time)
	if err != nil {
		t.Errorf("Expected valid RFC3339 time format, got %s", responseData.Time)
	}

	// Check that the time is recent (within last 5 seconds)
	now := time.Now().UTC()
	timeDiff := now.Sub(parsedTime)
	if timeDiff < 0 || timeDiff > 5*time.Second {
		t.Errorf("Expected recent time, got %s (diff: %v)", responseData.Time, timeDiff)
	}

	// Verify the response structure
	if responseData.Time == "" {
		t.Error("Expected non-empty time in response")
	}
}
