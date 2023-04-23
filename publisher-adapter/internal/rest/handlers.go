package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		errorResponse(w, "Content type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	_, err = http.Post("http://localhost:8081", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	//defer resp.Body.Close()
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
