FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

# Собираем сервис
WORKDIR /app/cmd/auth
RUN go build -o auth_service

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/cmd/auth/auth_service .
COPY --from=builder /app/config/config.yaml ./config/config.yaml
COPY --from=builder /app/db/migrations /root/db/migrations

EXPOSE 8080

CMD ["./auth_service"]
