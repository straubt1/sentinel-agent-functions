package main

import (
	"encoding/json"
	"net/http"
	"strings"

	xj "github.com/basgys/goxml2json"
)

// The inputs from Sentinel to the agent function
type XmlToJsonInputs struct {
	Content string `json:"content"`
}

// The outputs from the agent function to Sentinel
type XmlToJsonOutputs struct {
	Json string `json:"json"`
}

func XmlToJson(w http.ResponseWriter, r *http.Request) {
	data := &XmlToJsonInputs{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", ERROR_STATUS)
		return
	}
	xml := strings.NewReader(data.Content)
	result, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	resp := XmlToJsonOutputs{
		Json: result.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
