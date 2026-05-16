# Go Bookstore API

A modern, industry-level bookstore management API built with Go, GORM, and Uber Fx.

## Features

- **Layered Architecture**: Controller-Service-Repository pattern.
- **Dependency Injection**: Automated wiring using **Uber Fx**.
- **Authentication**: JWT-based authentication with `bcrypt` password hashing.
- **RBAC**: Role-Based Access Control with support for custom permissions.
- **Fluent Routing**: Laravel-inspired routing API for clean route definitions.
- **API Documentation**: Automated **Swagger** documentation (Swagger UI).
- **Database**: Automated migrations and structured seeders.
- **Validation**: Robust request validation using `go-playground/validator`.
- **Logging**: File-based error logging (`logs/error.log`).
- **Configuration**: Environment variable management via `.env`.

## Tech Stack

- **Go 1.26+**
- **Framework**: Gorilla Mux (with custom fluent wrapper)
- **ORM**: GORM (MySQL)
- **DI**: Uber Fx
- **Auth**: JWT-v5, Bcrypt
- **Docs**: Swaggo
- **Testing**: Testify

## Getting Started

### 1. Prerequisites
- Go installed on your machine.
- MySQL server running.

### 2. Configuration
Create a `.env` file in the root directory (one has been provided for you):
```env
DB_USER=root
DB_PASS=yourpassword
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=go_bookstore
SERVER_PORT=9010
JWT_SECRET=yoursecretkey
```

### 3. Installation
```bash
go mod tidy
```

### 4. Running the Application
Using the provided Makefile:
```bash
make run
```

### 5. Database Management
Like Laravel, you can refresh your database and seed it using CLI flags:
```bash
# Refresh database (Drop all tables and Migrate)
make migrate-fresh

# Refresh database and Seed initial data
make migrate-fresh-seed
```

### 6. API Documentation
The application automatically generates Swagger documentation. Once the server is running, you can access the interactive UI at:
`http://localhost:9010/swagger/index.html`

### 7. Testing
```bash
make test
```

## API Endpoints

### Auth
- `POST /register`: Register a new user.
- `POST /login`: Login and receive a JWT.

### Books (Protected)
- `GET /book`: Get all books.
- `POST /book`: Create a new book.
- `GET /book/{id}`: Get a specific book.
- `PUT /book/{id}`: Update a book (**Admin Only**).
- `DELETE /book/{id}`: Delete a book (**Admin Only**).

## Project Structure

- `cmd/app`: Entry point and application wiring.
- `pkg/config`: Configuration management.
- `pkg/controllers`: HTTP handlers.
- `pkg/database`: DB connection, migrations, and seeders.
- `pkg/middleware`: Auth, Logging, and RBAC middlewares.
- `pkg/models`: Data structures and validation tags.
- `pkg/repository`: Data access layer.
- `pkg/responses`: Standardized API responses.
- `pkg/routes`: Fluent routing definitions.
- `pkg/services`: Business logic layer.
- `pkg/validators`: Request validation logic.
