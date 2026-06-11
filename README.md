# Ainyx Solutions - Go Backend Task

A high-performance RESTful API built in Go to manage users and dynamically calculate their age.

## 🏗️ Architecture & Stack
- **GoFiber**: High-performance HTTP framework (Express.js equivalent).
- **SQLC**: Generates fully type-safe Go code from raw SQL queries.
- **PostgreSQL**: Relational database.
- **Uber Zap**: High-speed, structured JSON logging.
- **go-playground/validator**: Struct-based request validation.
- **Docker & Docker Compose**: Containerized infrastructure for instant setup.

## 🚀 How to Run

You do not need Go installed on your machine. You only need Docker.

1. Clone the repository and navigate into it:
   ```bash
   git clone <your-repo-url>
   cd ainyx-go-backend
   ```

2. Start the database and API simultaneously:
   ```bash
   docker-compose up --build
   ```

The API will be available at `http://localhost:3000`.

## 🧪 Running Unit Tests

To run the test suite for the dynamic age calculation logic:
```bash
go test ./internal/service -v
```

## 📡 API Endpoints

### 1. Create a User
```bash
curl -X POST http://localhost:3000/users \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice", "dob": "1990-05-10"}'
```

### 2. Get User by ID (Dynamic Age Calculation)
```bash
curl http://localhost:3000/users/1
```

### 3. List All Users (Paginated)
```bash
curl "http://localhost:3000/users?limit=10&offset=0"
```

### 4. Delete a User
```bash
curl -X DELETE http://localhost:3000/users/1
```

## ✨ Bonus Features Implemented
- **Dockerized Environment**: Single command startup.
- **Request Duration Logging**: Custom middleware logging exact ms duration via Zap.
- **Request ID Injection**: Custom middleware for tracing requests.
- **Pagination**: Implemented `LIMIT` and `OFFSET` on the GET `/users` endpoint.
- **Unit Testing**: Table-driven tests validating leap years and upcoming birthdays in the age calculator.
