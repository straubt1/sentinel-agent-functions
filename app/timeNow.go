package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// The inputs from Sentinel to the agent function
type TimeNowInputs struct {
	Format string `json:"format"`
}

// The outputs from the agent function to Sentinel
type TimeNowOutputs struct {
	Time string `json:"time"`
}

func TimeNow(w http.ResponseWriter, r *http.Request) {
	data := &TimeNowInputs{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", ERROR_STATUS)
		return
	}

	// check for specific formats or if non-empty, else use default
	format := time.RFC3339 //default
	if data.Format == "RFC3339" {
		format = time.RFC3339
	} else if data.Format == "RFC3339Nano" {
		format = time.RFC3339Nano
	} else if data.Format != "" {
		format = data.Format
	}

	resp := TimeNowOutputs{
		Time: time.Now().UTC().Format(format),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
