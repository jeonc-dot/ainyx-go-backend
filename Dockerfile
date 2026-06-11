# Build stage: We use the official Golang image to compile the code
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Compile the binary without C dependencies for maximum portability
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# Run stage: We use an ultra-lightweight alpine image to actually run it
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 3000
CMD ["./server"]
