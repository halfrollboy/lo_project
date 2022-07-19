package repository

import (
	"database/sql"
	"fmt"
	"lo_project/internal/domain/model"
	"log"
)

type PostgresDB interface {
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
}

type storage struct {
	client PostgresDB
}

func NewPgStorage(client PostgresDB) *storage {
	return &storage{client: client}
}

//Вспомогательная ф-ция для получения данных из model.json
func (s storage) CreateOrder(order model.Order) (err error) {
	_, err = s.client.Exec("INSERT INTO orders (order_log) VALUES($1)", order)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//Получаем элемент из кэша по ID
func (s storage) GetOrder(id string) (model.Order, error) {
	item := new(model.Item)
	err := s.client.QueryRow("SELECT id, order_log FROM orders ORDER BY id DESC LIMIT 1").Scan(&item.ID, &item.Orders)
	if err != nil {
		return model.Order{}, err
	}
	//отладочный принт
	fmt.Println("ItemID: %d, UUID: %s", item.ID, item.Orders.OrderUID)
	return item.Orders, nil
}
