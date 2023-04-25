package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		errorResponse(w, "Content type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	fmt.Println(r.Body)

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-type", "application/json").
		SetBody(r.Body).
		Post("http://localhost:7071")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Body())
	_, err = w.Write(resp.Body())
	if err != nil {
		log.Fatal(err)
	}

}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
