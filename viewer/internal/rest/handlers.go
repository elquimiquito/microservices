package rest

import (
	"Wildberries_services/repository"
	"Wildberries_services/repository/proxy"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var e repository.Employee
	body, err := io.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	err = json.Unmarshal(body, &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.ID)
	employee := proxy.GetByID(e.ID)
	data, err := json.Marshal(employee)
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
