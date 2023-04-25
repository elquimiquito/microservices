package reader

import (
	"Wildberries_services/reader/internal/nats"
	"Wildberries_services/repository"
	"Wildberries_services/repository/proxy"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

func Reader() {
	sc := nats.OpenNATS()
	sc.Subscribe("msg", insert, stan.StartAtTime(time.Now()), stan.SetManualAckMode())
	defer sc.Close()
	for 1 > 0 {

	}
}

func insert(m *stan.Msg) {
	err := m.Ack()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Привет")
	var e repository.Employee
	err = json.Unmarshal(m.Data, &e)
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	proxy.Insert(&e)
}
