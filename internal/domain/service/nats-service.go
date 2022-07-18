package service

import (
	"encoding/json"
	"lo_project/internal/domain/model"
	"lo_project/pkg/utils"
	"log"

	"github.com/nats-io/nats.go"
)

type Brocker interface {
	QueueSubscribe(subj, queue string, cb nats.MsgHandler) (*nats.Subscription, error)
	Publish(subj string, data []byte) error
}

type brokerStruct struct {
	broker Brocker
}

// Честно не знал как лучше архитектурно сделать с брокером поэтому убрал в сервис
func (b brokerStruct) Listen() (data model.Order) {
	b.broker.QueueSubscribe("foo", "queue.foo", func(msg *nats.Msg) {

		//Полуяаем дату
		log.Println("Subscriber 1:", string(msg.Data))
		err := json.Unmarshal([]byte(msg.Data), &data)
		if err != nil {
			log.Println("Ошибка отркытия файла")
		}
		//Пищем в базу

	})
	return data
}

// Сделать запись в бд
func (b brokerStruct) Write() {
	message := utils.GetDataModel()
	data, _ := json.Marshal(message)

	if err := b.broker.Publish("foo", []byte(data)); err != nil {
		log.Fatalln(err)
	}
}

// time.Sleep(2 * time.Second)
