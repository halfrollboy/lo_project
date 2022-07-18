package config

import (
	"fmt"
	"os"
)

type PgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func (pg PgConfig) GetDatabaseUrl() string {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.User, pg.Password, pg.Dbname)
	return psqlInfo
}

// Проверить задание конфига через структуру
//TODO Добавить обработчик ошибок на переменные
func NewConfig() *PgConfig {
	user, _ := os.LookupEnv("PG_USERNAME")
	password, _ := os.LookupEnv("PG_PASSWORD")
	dbname, _ := os.LookupEnv("PG_DATABASE")
	host, _ := os.LookupEnv("PG_HOST")
	port, _ := os.LookupEnv("GITHUB_USERNAME")
	return &PgConfig{Host: host, User: user, Password: password, Port: port, Dbname: dbname}
}
