package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // Import pq for side effects, such as registering its driver.
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	const (
		host   = "localhost"
		port   = 5432 // Default PostgreSQL port
		user   = "service_account"
		dbname = "cookbook"
	)
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, dbPassword, dbname)

	// Open a connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	http.HandleFunc("/", HelloServer)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This should automatically update on the server")
}
