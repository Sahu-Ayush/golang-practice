package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read custom headers
	requestID := r.Header.Get("X-Request-ID")
	userRole := r.Header.Get("X-User-Role")

	// Read standard headers
	userAgent := r.Header.Get("User-Agent")

	log.Printf("Request received | requestID=%s role=%s agent=%s",
		requestID, userRole, userAgent)

	// Respond
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{
  "message": "Hello from Go server",
  "request_id": "%s",
  "role": "%s"
}`, requestID, userRole)
}

func main() {
	http.HandleFunc("/hello", handler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
