package web_server

import (
	"bytes"
	"fmt"
	"github.com/karlseguin/gerb"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"
)

type WebServer struct {
	ContentDirectory string
}

func (s WebServer) getEnvironmentMap() map[string]interface{} {
	var envmap = make(map[string]interface{})
	for _, kv := range os.Environ() {
		parts := strings.Split(strings.ToLower(kv), "=")
		envmap[parts[0]] = parts[1]
	}
	envmap["hostname"], _ = os.Hostname()
	envmap["random"] = rand.Intn(99999)
	return envmap
}

func (s WebServer) webHandler(w http.ResponseWriter, r *http.Request) {
	template, err := gerb.ParseFile(false, fmt.Sprintf("%s/%s.gerb", s.ContentDirectory, path.Base(r.URL.Path)))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buffer := new(bytes.Buffer)
	template.Render(buffer, s.getEnvironmentMap())
	fmt.Fprintln(w, buffer)
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

func (s *WebServer) Run(port int, contentDirectory string) {
	s.ContentDirectory = contentDirectory
	http.HandleFunc("/web/", s.webHandler)
	http.HandleFunc("/health", s.HealthCheckHandler)
	log.Printf("Starting web server (content dir: %v) using port: %v", contentDirectory, port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
