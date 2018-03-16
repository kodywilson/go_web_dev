package main

import (
	"log"
	"net/http"
)

func main() {
	// PORT := ":8080"
	// log.Print("Running server on " + PORT)
	// http.HandleFunc("/", CompleteTaskFunc)
	// log.Fatal(http.ListenAndServe(PORT, nil))
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
