package main

// TODO: Better organize services

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tylorkolbeck/go-cookbook/api/v1/handlers"
	"github.com/tylorkolbeck/go-cookbook/auth"
	"github.com/tylorkolbeck/go-cookbook/internal/db"

	"github.com/tylorkolbeck/go-cookbook/internal/repository/cookbookRepo"
	"github.com/tylorkolbeck/go-cookbook/internal/repository/endpointRepo"
	"github.com/tylorkolbeck/go-cookbook/internal/repository/recipeRepo"
	"github.com/tylorkolbeck/go-cookbook/internal/repository/userRepo"
	"github.com/tylorkolbeck/go-cookbook/internal/service/cookbook"
	"github.com/tylorkolbeck/go-cookbook/internal/service/endpoints"
	"github.com/tylorkolbeck/go-cookbook/internal/service/recipe"
	"github.com/tylorkolbeck/go-cookbook/internal/service/user"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import pq for side effects, such as registering its driver.
)

func main() {
	dbConn := ConnectToDb()

	// Migrate the schema
	dbErr := db.AutoMigrate(dbConn)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new router
	router := gin.Default()

	// // Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Set the router to serve static files
	router.Static("/static", "./public/playground")

	// Define routes
	authConfig := auth.NewAuthConfig([]byte(os.Getenv("TOKEN_SECRET")))

	// USERS
	userRepository := userRepo.NewPostgresUserRepository(dbConn)
	userService := user.Initialize(userRepository, *authConfig)
	handlers.RegisterUserRoutes(router, userService, authConfig)

	// COOKBOOKS
	cookbookRepo := cookbookRepo.NewPostgresCookbookRepository(dbConn)
	cookbookService := cookbook.Initialize(cookbookRepo)
	handlers.RegisterCookbookRoutes(router, cookbookService)

	// RECIPES
	recipeRepo := recipeRepo.NewPostgresRecipeRepository(dbConn)
	recipeService := recipe.Initialize(recipeRepo)
	handlers.RegisterRecipeRoutes(router, recipeService)

	// ENDPOINTS
	endpointRepo := endpointRepo.NewEndpointRepository()
	endpointService := endpoints.Initialize(*endpointRepo)
	endpointHandler := handlers.NewEndpointsHandler(endpointService)

	// ENDPOINTS
	router.GET("/endpoints", endpointHandler.ListEndpoints)
	router.StaticFile("/playground", "./public/playground/client/src/index.html")
	// Apply CORS middleware

	// Start the server
	router.Run(":8080")
}

func ConnectToDb() *gorm.DB {
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"),
	// )

	dsn := "host=localhost user=postgres password=password dbname=cookbook port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
		panic("failed	to	connect database")
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	fmt.Println("Successfully connected to the DB!")

	return db
}
