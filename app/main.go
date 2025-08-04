package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{ID: 1, Title: "First Post", Body: "This is the first post"},
		{ID: 2, Title: "Second Post", Body: "This is the second post"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func envHandler(w http.ResponseWriter, r *http.Request) {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			env[pair[0]] = pair[1]
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(env)
}

// func getAmi(w http.ResponseWriter, r *http.Request) {
// 	ami := os.Getenv("AWS_AMI_ID")
// 	if ami == "" {
// 		http.Error(w, "AMI ID not set", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{"ami_id": ami})
// }

func main() {
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/envs", envHandler)

	fmt.Println("Server starting on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
