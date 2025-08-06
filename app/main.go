package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// convert XML to JSON
// Parse HCL
// substring

// Single error status code for all functions (keep it simple)
const ERROR_STATUS = http.StatusInternalServerError

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/timeNow", TimeNow)
	r.Post("/parseHCL", ParseHCL)
	r.Post("/xmlToJson", XmlToJson)

	// Port configuration
	port := 5000
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic("Invalid port number")
		}
	}
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
