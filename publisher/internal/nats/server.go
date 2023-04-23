package nats

import (
	"github.com/nats-io/stan.go"
	"log"
)

func openNATS() stan.Conn {
	sc, err := stan.Connect("test-cluster", "publisher")
	if err != nil {
		log.Printf("Ошибка %s", err)
	}
	return sc
}
