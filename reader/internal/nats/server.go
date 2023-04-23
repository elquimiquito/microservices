package nats

import (
	"github.com/nats-io/stan.go"
	"log"
)

func OpenNATS() stan.Conn {
	sc, err := stan.Connect("test-cluster", "reader")
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	return sc
}
