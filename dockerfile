# Build Stage
FROM golang:1.24.1-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY .env ./
COPY . .

RUN go build -o app main.go

# Run Stage
FROM debian:bookworm-slim

WORKDIR /app

# Copy both the app binary and .env file
COPY --from=builder /app/app .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./app"]

