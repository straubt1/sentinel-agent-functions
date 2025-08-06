package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseHCL(t *testing.T) {
	// Test data - sample HCL content
	testData := ParseHCLInputs{
		Content: `
resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"
  
  tags = {
    Name = "example-instance"
  }
}`,
	}

	// Convert test data to JSON
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	// Create request
	req := httptest.NewRequest("POST", "/parseHCL", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Execute the ParseHCL function directly
	ParseHCL(w, req)

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
	var responseData ParseHCLOutputs
	if err := json.NewDecoder(w.Body).Decode(&responseData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Verify the response structure
	if responseData.Json == "" {
		t.Error("Expected non-empty JSON in response")
	}

	// Verify that the response contains valid JSON
	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(responseData.Json), &parsedJSON); err != nil {
		t.Errorf("Expected valid JSON in response, got: %s, error: %v", responseData.Json, err)
	}

	// For the current implementation, verify the expected static response
	expectedJSON := `{"parsed": true}`
	if responseData.Json != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, responseData.Json)
	}
}
