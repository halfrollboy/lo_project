package main

import (
	"lo_project/internal/api/handler"
	"lo_project/internal/config"
	"log"
)

func main() {
	if err := startApp(); err != nil {
		log.Fatal("Приложение не стартовало")
	}
}

//Ф-ция для обработки ошибки запуска приложения
func startApp() error {
	log.Println("main start")

	appConfig := config.NewAppConfig()
	router := handler.NewRouter(appConfig) //Наверно не стоит так передавать
	router.Router.Run()
	return nil
}
