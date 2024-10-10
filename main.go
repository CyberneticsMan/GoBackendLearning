package main

import (
	"net/http"

	"github.com/CyberneticsMan/GoBackendLearning/config"
	"github.com/CyberneticsMan/GoBackendLearning/controller"
	_ "github.com/CyberneticsMan/GoBackendLearning/docs"
	"github.com/CyberneticsMan/GoBackendLearning/helper"
	"github.com/CyberneticsMan/GoBackendLearning/model"
	"github.com/CyberneticsMan/GoBackendLearning/repository"
	"github.com/CyberneticsMan/GoBackendLearning/router"
	"github.com/CyberneticsMan/GoBackendLearning/service"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
