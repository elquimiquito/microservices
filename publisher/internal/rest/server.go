package rest

import "net/http"

func OpenHTTP() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	http.ListenAndServe("localhost:8081", mux)
}
