package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tylorkolbeck/go-cookbook/api/v1/handlers"
	"github.com/tylorkolbeck/go-cookbook/auth"
	"github.com/tylorkolbeck/go-cookbook/internal/repository"
	"github.com/tylorkolbeck/go-cookbook/internal/service"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // Import pq for side effects, such as registering its driver.
)

func main() {
	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new router
	router := gin.Default()

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
	// USERS
	// Initialize Auth
	authConfig := auth.NewAuthConfig([]byte(os.Getenv("TOKEN_SECRET")))
	userRepository := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepository, *authConfig)
	userHandler := handlers.NewUserHandler(userService)
	router.POST("/users", userHandler.CreateUser)
	router.POST("/login", func(c *gin.Context) {
		userHandler.Login(c, authConfig)
	})
	router.GET("/users", userHandler.ListUsers)
	router.GET("/verify/:token", userHandler.VerifyEmail)
	// COOKBOOKS
	cookbookRepo := repository.NewInMemoryCookbookRepository()
	cookbookService := service.NewCookbookService(cookbookRepo)
	cookbookHandler := handlers.NewCookbookHandler(cookbookService)
	router.GET("/cookbooks", cookbookHandler.ListCookbooks)
	router.POST("/cookbooks", cookbookHandler.CreateCookbook)
	router.GET("/cookbooks/:id", cookbookHandler.GetCookbook)
	router.PUT("/cookbooks/:id", cookbookHandler.UpdateCookbook)
	router.DELETE("/cookbooks/:id", cookbookHandler.DeleteCookbook)

	// RECIPES
	recipeRepo := repository.NewInMemoryRecipeRepository()
	recipeService := service.NewRecipeService(recipeRepo)
	recipeHandler := handlers.NewRecipeHandler(recipeService)

	// RECIPES
	router.GET("/recipes", recipeHandler.ListRecipes)
	router.POST("/recipes", recipeHandler.CreateRecipe)
	router.GET("/recipes/:id", recipeHandler.GetRecipe)
	router.PUT("/recipes/:id", recipeHandler.UpdateRecipe)
	router.DELETE("/recipes/:id", recipeHandler.DeleteRecipe)

	// ENDPOINTS
	endpointRepo := repository.NewEndpointRepository()
	endpointService := service.NewEndpointsService(*endpointRepo)
	endpointHandler := handlers.NewEndpointsHandler(endpointService)

	// ENDPOINTS
	router.GET("/endpoints", endpointHandler.ListEndpoints)
	router.StaticFile("/playground", "./public/playground/client/src/index.html")
	// Apply CORS middleware

	// Start the server
	router.Run(":8080")

	// ConnectToDb()
	// StartHttpServer()
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This should automatically update on the server v1")
}

func StartHttpServer() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("The server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func ConnectToDb() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// Open a connection
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("Error: Could not establish a connection with the database")
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the DB!")
}
