package web_server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type WebServer struct {
}

func (s WebServer) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(100)
	if random%2 == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintln(w, fmt.Sprintf("Random: %d", random))
}

func (s WebServer) Run(port int, staticPagesdirectory string) {
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir(staticPagesdirectory))))
	http.HandleFunc("/health", s.HealthCheckHandler)
	log.Printf("Starting web server (static web dir: %v) using port: %v", staticPagesdirectory, port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
