# Go CRUD API

A simple RESTful CRUD API built with Go, using Gorilla Mux for routing and GORM with SQLite for data persistence.

## Features

- Create, Read, Update, and Delete operations for books
- SQLite database for data storage
- RESTful API endpoints
- JSON request/response format

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
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{"title":"The Go Programming Language","author":"Alan A. A. Donovan","isbn":"978-0134190440"}'
```
