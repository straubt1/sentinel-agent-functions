package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	// Test data
	testData := SleepInputs{
		DurationMilliseconds: 100, // Test with 100ms sleep
	}

	// Convert test data to JSON
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	// Create request
	req := httptest.NewRequest("POST", "/sleep", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Record start time to measure actual sleep duration
	startTime := time.Now()

	// Execute the Sleep function directly
	Sleep(w, req)

	// Record end time to verify sleep duration
	endTime := time.Now()
	actualDuration := endTime.Sub(startTime)

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
	var responseData SleepOutputs
	if err := json.NewDecoder(w.Body).Decode(&responseData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Verify the response message contains the expected duration
	expectedMessage := "Slept for 100 milliseconds"
	if responseData.Message != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, responseData.Message)
	}

	// Verify that the function actually slept for approximately the requested duration
	// Allow some tolerance for execution overhead (±50ms)
	expectedDuration := 100 * time.Millisecond
	minDuration := expectedDuration - 50*time.Millisecond
	maxDuration := expectedDuration + 50*time.Millisecond

	if actualDuration < minDuration || actualDuration > maxDuration {
		t.Errorf("Expected sleep duration between %v and %v, got %v", minDuration, maxDuration, actualDuration)
	}

	// Verify the response structure
	if responseData.Message == "" {
		t.Error("Expected non-empty message in response")
	}
}
