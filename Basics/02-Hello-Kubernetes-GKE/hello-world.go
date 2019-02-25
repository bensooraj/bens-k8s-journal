package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// SomeResponse .
type SomeResponse struct {
	HostName    string    `json:"hostname"`
	CurrentTime time.Time `json:"current_time"`
	Message     string    `json:"message"`
	ServicePort string    `json:"service_port"`
}

func main() {
	// Server Mux
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8000",
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	servicePort := os.Getenv("HELLO_NODE_SERVICE_PORT")
	data := SomeResponse{
		HostName:    hostname,
		CurrentTime: time.Now(),
		Message:     "Hello, World!",
		ServicePort: servicePort,
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
