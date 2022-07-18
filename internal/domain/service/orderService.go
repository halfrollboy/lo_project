package service

import "lo_project/internal/domain/model"

type CacheRepository interface {
	SetOrder(order model.Order) error
	GetOrder(id string) (model.Order, bool)
}

type PgRepository interface {
	CreateOrder() (err error)
	GetOrder(id string) (model.Order, error)
}

type orderService struct {
	cache CacheRepository
	db    PgRepository
}

func (os orderService) GetOrder(id string) (model.Order, error) {
	ordCach, found := os.cache.GetOrder(id)
	if !found {
		ordDb, err := os.db.GetOrder(id)
		if err != nil {
			return model.Order{}, err
		}
		return ordDb, nil
	}
	return ordCach, nil
}

func (os orderService) SetOrder(order model.Order) {
	os.db.CreateOrder(order)
}
