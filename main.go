package main

import "net/http"

func main() {
	http.HandleFunc("/", saludar)
	http.ListenAndServe(":8080", nil)
}

func saludar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
