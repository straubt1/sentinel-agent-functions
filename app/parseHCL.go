package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
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
	data := &ParseHCLInputs{}
	err := json.NewDecoder(r.Body).Decode(&data)

	b := []byte(data.Content)
	parser := hclparse.NewParser()
	f, parseDiags := parser.ParseHCL(b, "test.hcl")
	if parseDiags.HasErrors() {
		log.Fatal(parseDiags.Error())
	}

	// var fooInstance foo
	var result map[string]interface{}
	_ = gohcl.DecodeBody(f.Body, nil, &result)
	// if decodeDiags.HasErrors() {
	// 	log.Fatal(decodeDiags.Error())
	// }

	// fmt.Printf("%#v", fooInstance)

	if err != nil {
		http.Error(w, "Invalid HCL", ERROR_STATUS)
		return
	}
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to marshal parsed HCL", ERROR_STATUS)
		return
	}
	resp := ParseHCLOutputs{
		Json: string(jsonBytes),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
