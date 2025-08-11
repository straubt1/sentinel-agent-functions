package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// The inputs from Sentinel to the agent function
type SleepInputs struct {
	DurationMilliseconds int `json:"duration"`
}

// The outputs from the agent function to Sentinel
type SleepOutputs struct {
	Message string `json:"message"`
}

func Sleep(w http.ResponseWriter, r *http.Request) {
	data := &SleepInputs{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", ERROR_STATUS)
		return
	}
	// Sleep for the specified duration
	time.Sleep(time.Duration(data.DurationMilliseconds) * time.Millisecond)
	resp := SleepOutputs{
		Message: "Slept for " + strconv.Itoa(data.DurationMilliseconds) + " milliseconds",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
