package service

import (
	"lo_project/internal/domain/model"
	"log"
)

type CacheRepository interface {
	SetOrder(order model.Order) error
	GetOrder(id string) (model.Order, bool)
}

type PgRepository interface {
	CreateOrder(order model.Order) (err error)
	GetOrder(id string) (model.Order, error)
}

type orderService struct {
	cache CacheRepository
	db    PgRepository
}

func NewOrderService(cache CacheRepository, db PgRepository) *orderService {
	return &orderService{cache: cache, db: db}
}

func (os orderService) GetOrder(id string) (model.Order, error) {
	ordCach, found := os.cache.GetOrder(id)
	if !found {
		ordDb, err := os.db.GetOrder(id)
		if err != nil {
			return model.Order{}, err
		}
		e := os.cache.SetOrder(ordDb)
		if e != nil {
			log.Fatal("Не получилось записать order в кэш")
		}
		return ordDb, nil
	}
	return ordCach, nil
}

func (os orderService) SetOrder(order model.Order) error {
	err := os.db.CreateOrder(order)
	if err != nil {
		log.Fatal("Ошибка записи Order в базу")
		return err
	}
	e := os.cache.SetOrder(order)
	if e != nil {
		log.Fatal("Не получилось записать order в кэш")
		return e
	}
	return nil
}
