package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Heyyy"))
}

func sayPong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Custon-Header", "pingpong")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Pong"))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/ping", sayPong)
	http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux)
}
