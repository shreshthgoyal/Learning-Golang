package main

import (
	"bootcamp/pkg/api"
	"bootcamp/pkg/store"
	"log"
	"net/http"
)

func main() {
	str, err := store.NewLocal("db.gob")
	if err != nil {
		log.Fatal(err)
	}

	apy := api.New(str)

	mux := http.NewServeMux()

	apy.Register(mux)

	http.ListenAndServe(":8080", mux)
}
