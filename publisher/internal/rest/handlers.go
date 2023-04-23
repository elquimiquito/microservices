package rest

import (
	"Wildberries_services/publisher/internal/nats"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	nats.InitPublisher(data)
}
