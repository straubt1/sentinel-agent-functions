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

		processes, err := os.ReadDir("/proc")
		if err != nil {
			http.Error(w, "Failed to read /proc", ERROR_STATUS)
			return
		}
		w.Write([]byte("\nAvailable processes:\n"))
		for _, entry := range processes {
			w.Write([]byte(entry.Name() + "\n"))

			println(entry.Name())
			// if entry.IsDir() {
			// 	// /proc directories with numeric names are processes
			// 	if _, err := strconv.Atoi(entry.Name()); err == nil {
			// 		w.Write([]byte(entry.Name() + "\n"))
			// 	}
			// }
		}
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
	println("Starting server on port", port)
	http.ListenAndServe(":"+strconv.Itoa(port), r)
	// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout localhost.key -out localhost.crt -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
	// http.ListenAndServeTLS(":"+strconv.Itoa(port), "localhost.crt", "localhost.key", r)
}
