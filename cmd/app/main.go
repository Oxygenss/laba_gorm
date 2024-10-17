package main

import (
	"log"

	"github.com/Oxygenss/laba_gorm/internal/handler"
	"github.com/Oxygenss/laba_gorm/internal/repository"
)

func main() {
	_, err := repository.Init()
	if err != nil {
		log.Fatalf("failed to init database: %v", err)
	}

	router := handler.InitRoutes()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
