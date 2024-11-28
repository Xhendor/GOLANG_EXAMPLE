# Go CRUD API

A simple RESTful CRUD API built with Go, using Gorilla Mux for routing and GORM with SQLite for data persistence.

# Project Structure:
- models/: Contains the Book model and error definitions
- handlers/: Contains HTTP request handlers
- services/: Contains business logic
- middleware/: Contains HTTP middleware (logging)
- config/: Contains configuration management


## Prerequisites

- Go 1.21 or higher
- SQLite

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```

## Running the Application
## Default Credentials:
- Username: admin
- Password: password

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

- `GET /books` - Get all books
- `GET /books/{id}` - Get a specific book
- `POST /books` - Create a new book
- `PUT /books/{id}` - Update a book
- `DELETE /books/{id}` - Delete a book

## Example Request

Create a new book:
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'

curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json " \
-H "Authorization: Bearer YOUR_TOKEN" \
-d '{"title":"The Go Programming Language","author":"Alan A. A. Donovan","isbn":"978-0134190440"}'
```
# In Swagger UI (http://localhost:8080/swagger/):

- Click the "Authorize" button at the top
- Enter your token in the format: Bearer YOUR_TOKEN
- Click "Authorize"
- Now you can test all protected endpoints
- The authentication system includes:

- Login endpoint (/login)
- JWT token generation
- Token validation middleware
- Protected routes for all book operations
- Swagger documentation with authentication


