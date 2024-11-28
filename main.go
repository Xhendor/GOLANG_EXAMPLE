package main

import (
	"log"
	"net/http"

	"go-crud/config"
	"go-crud/handlers"
	"go-crud/middleware"
	"go-crud/models"
	"go-crud/services"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&models.Book{})

	// Initialize services and handlers
	bookService := services.NewBookService(db)
	bookHandler := handlers.NewBookHandler(bookService)

	// Initialize router
	r := mux.NewRouter()

	// Add middleware
	r.Use(middleware.Logging)

	// Define routes
	r.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	// Start server
	log.Printf("Server is running on http://localhost:%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
