FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service cmd/main.go

FROM alpine:latest

COPY --from=builder /app/auth-service /auth-service

EXPOSE 3000

CMD ["/auth-service"]
