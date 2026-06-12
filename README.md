# Ainyx Go Backend Task

This is my submission for the Ainyx backend engineering task. It's a simple Go microservice that handles user creation and dynamic age calculation.

I decided to over-engineer the architecture just a little bit to ensure it was production-ready.

## Tech Stack & Design Choices
- **GoFiber**: Used this instead of standard `net/http` for better routing and built-in middleware support.
- **SQLC & PostgreSQL**: I explicitly avoided using an ORM (like GORM) because of the reflection overhead. SQLC generates type-safe Go code straight from my raw SQL queries, which keeps the memory footprint tiny and prevents SQL injection.
- **Uber Zap**: Used for structured, high-speed JSON logging.
- **Docker**: Containerized both the Go app and the Postgres DB so you can spin it up with one command.

## How to run it

You don't need Go installed locally to test this, just Docker.

1. Clone it and cd into the directory:
```bash
git clone https://github.com/jeonc-dot/ainyx-go-backend.git
cd ainyx-go-backend
```

2. Spin up the containers:
```bash
docker-compose up --build
```
The API will be live on `http://localhost:3000`.

## API Endpoints

I implemented pagination as a bonus for the GET all users route.

**Create a User**
```bash
curl -X POST http://localhost:3000/users -H "Content-Type: application/json" -d '{"name": "Alice", "dob": "1990-05-10"}'
```

**Get User by ID (Calculates Age dynamically)**
```bash
curl http://localhost:3000/users/1
```

**List Users (Paginated)**
```bash
curl "http://localhost:3000/users?limit=10&offset=0"
```

**Delete User**
```bash
curl -X DELETE http://localhost:3000/users/1
```

## Testing

If you want to run the unit tests (which cover edge cases like leap years and upcoming birthdays in the age calculator):

```bash
go test ./internal/service -v
```

## Bonus stuff I added
- Docker compose setup.
- Custom middleware that tracks exact request duration (ms) using Zap.
- Request ID injection for tracing.
- Pagination for the list endpoint.
