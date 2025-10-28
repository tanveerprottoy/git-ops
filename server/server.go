package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tanveerprottoy/git-ops/mathext"
)

// mainHandler handles all HTTP methods for the endpoints
func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Set JSON response header
	w.Header().Set("Content-Type", "application/json")

	log.Printf("Received %s request for %s", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	case http.MethodPut:
		handlePut(w, r)
	case http.MethodDelete:
		handleDelete(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Printf("request header: %v\n", r.Header)

	msg := fmt.Sprintf("Hello, from GET Handler. Adding 2 + 3 = %d", mathext.Add(2, 3))
	log.Println(msg)

	i, err := w.Write([]byte(msg))
	log.Printf("i: %d, err: %s", i, err.Error())
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	log.Printf("request header: %v\n", r.Header)

	msg := fmt.Sprintf("Hello, from POST Handler. Adding 1 + 4 = %d", mathext.Add(1, 4))
	log.Println(msg)

	i, err := w.Write([]byte(msg))
	log.Printf("i: %d, err: %s", i, err.Error())
}

// PUT: Update item by ID (expects /{id} as path param)
func handlePut(w http.ResponseWriter, r *http.Request) {
	log.Printf("request header: %v\n", r.Header)

	msg := fmt.Sprintf("Hello, from PUT Handler. Adding 2 + 1 = %d", mathext.Add(2, 1))
	log.Println(msg)

	i, err := w.Write([]byte(msg))
	log.Printf("i: %d, err: %s", i, err.Error())
}

// DELETE: Remove item by ID (expects /{id} as path param)
func handleDelete(w http.ResponseWriter, r *http.Request) {
	log.Printf("request header: %v\n", r.Header)

	msg := fmt.Sprintf("Hello, from DELETE Handler. Adding 6 + 5 = %d", mathext.Add(6, 5))
	log.Println(msg)

	i, err := w.Write([]byte(msg))
	log.Printf("i: %d, err: %s", i, err.Error())
}

func Serve() {
	http.HandleFunc("/", mainHandler)

	fmt.Println("Server running at http://localhost:8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
