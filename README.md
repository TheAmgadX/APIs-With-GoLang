# Go API Project

This is a simple API built using Go and the Chi router. It follows a structured project layout for maintainability and scalability.

## Project Structure

```
.
├── api                   # API-related logic
├── internal              # Internal packages
│   ├── tools             # Utility functions and helpers
│   ├── handlers          # Request handlers for different endpoints
│   ├── middleware        # Middleware functions (e.g., authentication)
├── cmd                   # Entry point of the application
│   ├── api               # Main API command
├── go.mod                # Go module file
├── go.sum                # Dependency checksum file
```

## Features

- **Chi Router** for handling routes efficiently.
- **Middleware** for authentication and request processing.
- **Modular Structure** for better code organization.
- **Error Handling** with structured responses.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/TheAmgadX/APIs-With-GoLang.git
   cd APIs-With-GoLang
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Run the application:

   ```sh
   go run cmd/api/main.go
   ```

## API Endpoints

### Account Routes

| Method | Endpoint       | Description              |
|--------|--------------|--------------------------|
| GET    | `/account/coins` | Retrieves coin balance |

## Middleware

The project includes middleware for request validation and authentication.
