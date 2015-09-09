package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello webhook!")
}

func main() {
	http.Handle("/", http.HandlerFunc(indexHandler))
	http.Handle("/webhook", http.HandlerFunc(webhookHandler))

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	panic(http.ListenAndServe(":2003", nil))
}
