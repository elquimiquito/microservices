package rest

import "net/http"

func OpenHTTP() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	http.ListenAndServe("localhost:7071", mux)
}
