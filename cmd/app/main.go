package main

import (
	"github.com/Oxygenss/laba_gorm/internal/database"
	"github.com/Oxygenss/laba_gorm/internal/handler"
	"github.com/pkg/errors"
)

func main() {
	_, err := database.Init()
	if err != nil {
		errors.Wrap(err, "get all")
	}

	router := handler.InitRoutes()
	router.Run(":8081")
}
