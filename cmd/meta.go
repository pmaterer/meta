package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/pmaterer/meta/config"
	"github.com/pmaterer/meta/internal/postgres"
	"github.com/pmaterer/meta/slip/delivery/http"
	"github.com/pmaterer/meta/slip/repository"
	"github.com/pmaterer/meta/slip/service"
)

func main() {
	var config config.Config
	err := envconfig.Process("meta", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := postgres.NewHandler(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	slipRepo := repository.NewRepository(db)
	slipService := service.NewService(slipRepo)
	slipHandler := http.NewHandler(slipService)

	r := gin.Default()
	r.POST("/slips", slipHandler.CreateSlip)
	r.GET("/slips/:id", slipHandler.GetSlip)
	r.GET("/slips", slipHandler.GetAllSlips)
	r.PUT("/slips/:id", slipHandler.UpdateSlip)
	r.DELETE("/slips/:id", slipHandler.DeleteSlip)

	log.Fatal(r.Run(fmt.Sprintf("%s:%d", config.ServerListenAddress, config.ServerListenPort)))
}
