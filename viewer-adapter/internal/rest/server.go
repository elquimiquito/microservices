package rest

import (
	"log"
	"net/http"
)

func OpenHTTP() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	err := http.ListenAndServe("localhost:7080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
