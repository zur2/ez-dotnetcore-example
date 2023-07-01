// main.go
package main

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to the index.html file
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Failed to get current working directory", http.StatusInternalServerError)
		return
	}
	htmlPath := filepath.Join(dir, "index.html")

	html, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		http.Error(w, "Failed to read HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}
