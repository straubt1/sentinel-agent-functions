package main

import (
	"encoding/json"
	"net/http"
)

// The inputs from Sentinel to the agent function
type ParseHCLInputs struct {
	Content string `json:"content"`
}

// The outputs from the agent function to Sentinel
type ParseHCLOutputs struct {
	Json string `json:"json"`
}

func ParseHCL(w http.ResponseWriter, r *http.Request) {
	// data := &ParseHCLInputs{}
	// err := json.NewDecoder(r.Body).Decode(&data)
	// if err != nil {
	// 	http.Error(w, "Invalid JSON", ERROR_STATUS)
	// 	return
	// }
	resp := ParseHCLOutputs{
		Json: "{\"parsed\": true}",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
