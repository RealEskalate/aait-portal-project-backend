# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /main ./main.go

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /main /app/main
COPY .env /app/.env  

EXPOSE 8080
ENV GIN_MODE=release
CMD ["/app/main"]