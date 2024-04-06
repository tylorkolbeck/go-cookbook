package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tylorkolbeck/go-cookbook/api/v1/handlers"
	"github.com/tylorkolbeck/go-cookbook/auth"
	"github.com/tylorkolbeck/go-cookbook/internal/config"
	"github.com/tylorkolbeck/go-cookbook/internal/db"
	"github.com/tylorkolbeck/go-cookbook/internal/http"
	"github.com/tylorkolbeck/go-cookbook/internal/repository"
	"github.com/tylorkolbeck/go-cookbook/internal/service"

	_ "github.com/lib/pq"
)

type App struct {
	Router   *gin.Engine
	Services *service.AppServices
	Repos    *repository.AppRepository
}

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.ConnectToDb(config)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}

	authConfig, err := auth.NewAuthConfig(config.TOKEN_SECRET)
	if err != nil {
		log.Fatal("Could not create auth config", err)
	}

	router := http.SetupRouter()

	repos, err := repository.NewAppRepository(dbConn)
	if err != nil {
		log.Fatal("Could not create app repository", err)
	}

	services := service.NewAppServices(repos, *authConfig)

	handlers.RegisterRoutesAndHandlers(router, *services, authConfig)

	app := &App{
		Router:   router,
		Services: services,
		Repos:    &repos,
	}

	dbErr := db.AutoMigrate(dbConn)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	app.Router.Run(fmt.Sprintf(":%s", config.PORT))
}
