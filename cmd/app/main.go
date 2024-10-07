package main

import (
	"github.com/Oxygenss/laba_gorm/internal/handler"
	"github.com/Oxygenss/laba_gorm/internal/repository"
	"github.com/pkg/errors"
)

func main() {
	_, err := repository.Init()
	if err != nil {
		errors.Wrap(err, "failed to init database")
	}

	router := handler.InitRoutes()
	router.Run(":8080")
}
