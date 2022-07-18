package main

import (
	"fmt"
	"lo_project/internal/api/handler"
	"log"
)

func main() {
	if err := startApp(); err != nil {
		log.Fatal("Приложение не стартовало")
	}
}

func startApp() error {
	fmt.Println("main start")
	router := handler.NewRouter()
	router.Router.Run()
	return nil
}
