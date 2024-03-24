package main

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
	"github.com/tylorkolbeck/go-cookbook/api/v1/handlers/cookbookHandler"
	"github.com/tylorkolbeck/go-cookbook/auth"
	"github.com/tylorkolbeck/go-cookbook/internal/db"
	"github.com/tylorkolbeck/go-cookbook/internal/repository"
	"github.com/tylorkolbeck/go-cookbook/internal/repository/cookbookRepo"
	"github.com/tylorkolbeck/go-cookbook/internal/repository/recipeRepo"
	"github.com/tylorkolbeck/go-cookbook/internal/service/cookbook"
	"github.com/tylorkolbeck/go-cookbook/internal/service/endpoints"
	"github.com/tylorkolbeck/go-cookbook/internal/service/recipe"
	"github.com/tylorkolbeck/go-cookbook/internal/service/user"
	"github.com/tylorkolbeck/go-cookbook/middleware"

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
	userRepository := repository.NewInMemoryUserRepository()
	userService := user.Initialize(userRepository, *authConfig)
	userHandler := handlers.NewUserHandler(userService)

	router.POST("/users", userHandler.CreateUser)
	router.POST("/login", func(c *gin.Context) {
		userHandler.Login(c, authConfig)
	})
	router.GET("/users/:id", userHandler.GetUserByID)
	router.GET("/users", middleware.AuthMiddleware(), userHandler.ListUsers)
	router.GET("/verify/:token", userHandler.VerifyEmail)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	router.PUT("/users/:id", userHandler.UpdateUser)

	// COOKBOOKS
	cookbookRepo := cookbookRepo.NewPostgresCookbookRepository(dbConn)
	// cookbookRepo := cookbookRepo.NewInmemoryCookbookRepository()
	cookbookService := cookbook.Initialize(cookbookRepo)
	cookbookHandler.RegisterCookbookRoutes(router, cookbookService)

	// RECIPES
	recipeRepo := recipeRepo.NewInMemoryRecipeRepository()
	recipeService := recipe.Initialize(recipeRepo)
	recipeHandler := handlers.NewRecipeHandler(recipeService)
	router.GET("/recipes", recipeHandler.ListRecipes)
	router.POST("/recipes", recipeHandler.CreateRecipe)
	router.GET("/recipes/:id", recipeHandler.GetRecipe)
	router.PUT("/recipes/:id", recipeHandler.UpdateRecipe)
	router.DELETE("/recipes/:id", recipeHandler.DeleteRecipe)

	// ENDPOINTS
	endpointRepo := repository.NewEndpointRepository()
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
