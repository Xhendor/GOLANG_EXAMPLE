package main

import (
	"log"
	"net/http"

	"go-crud/config"
	_ "go-crud/docs"
	"go-crud/handlers"
	"go-crud/middleware"
	"go-crud/models"
	"go-crud/services"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title Book Store API
// @version 1.0
// @description This is a Book Store API server.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&models.Book{}, &models.User{})

	// Initialize services and handlers
	bookService := services.NewBookService(db)
	bookHandler := handlers.NewBookHandler(bookService)
	authHandler := handlers.NewAuthHandler()

	// Initialize router
	r := mux.NewRouter()

	// Add middleware
	r.Use(middleware.Logging)

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+cfg.ServerPort+"/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	// Public routes
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Protected routes
	api := r.PathPrefix("").Subrouter()
	api.Use(middleware.AuthMiddleware)

	// Book routes
	api.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	api.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	api.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
	api.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	api.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	// Start server
	log.Printf("Server is running on http://localhost:%s", cfg.ServerPort)
	log.Printf("Swagger documentation is available at http://localhost:%s/swagger/", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
