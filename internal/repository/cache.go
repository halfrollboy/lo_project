package repository

import (
	"fmt"
	"lo_project/internal/domain/model"
	"time"

	"github.com/patrickmn/go-cache"
)

type CaheInterface interface {
	Set(k string, x interface{}, d time.Duration)
	Get(k string) (interface{}, bool)
}

type cacheRepository struct {
	cache CaheInterface
}

//Вставляем по id в кэш стркутуру заказа
func (c cacheRepository) SetOrder(order model.Order) error {
	c.cache.Set(order.OrderUID, order, cache.DefaultExpiration)
	return nil
}

//Ищем по id в кэше если нет, то отдаём наверх false
func (c cacheRepository) GetOrder(id string) (model.Order, bool) {
	item, found := c.cache.Get(id)
	if !found {
		return model.Order{}, false
	}
	fmt.Println(item)
	return item.(model.Order), true
}
