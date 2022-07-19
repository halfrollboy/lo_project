package config

import (
	"lo_project/internal/domain/model"
	"lo_project/internal/domain/service"
	"lo_project/internal/repository"
	"lo_project/pkg/cash"
	natsstreaming "lo_project/pkg/client/nats-streaming"
	"lo_project/pkg/client/postgres"
	"log"
)

type PgService interface {
	SetOrder(order model.Order) error
	GetOrder(id string) (model.Order, error)
}

type NatsService interface {
	Listen() (data model.Order)
	Write()
}

type AppConfig struct {
	OrderService PgService
	NatsService  NatsService
}

//Стоит ли выносить все эти вызовы в вызов app
func NewAppConfig() *AppConfig {
	pg, err := postgres.CreateConn()
	if err != nil {
		log.Fatal("Не получил Conn базы")
	}
	return &AppConfig{
		// Подключаем сервис заказов
		OrderService: service.NewOrderService(
			//Подключаем репозиторий кэша
			repository.NewCacheReposytory(
				// Получаем коннект кэша
				cash.GetCacheConn(),
			),
			//Получаем репозиторий Postgres
			repository.NewPgStorage(
				//Передаём Conn от Postgres
				pg,
			),
		),
		// Получаем Nats сревис
		NatsService: service.NewBrockerStruct(
			natsstreaming.GetNatsConnect(),
		),
	}
}
