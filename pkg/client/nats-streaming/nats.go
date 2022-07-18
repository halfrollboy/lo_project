package natsstreaming

import (
	"log"

	"github.com/nats-io/nats.go"
)

//Понять как сделать через интерфейс
func GetNatsConnect() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	return nc
}
