FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY auth/go.mod auth/go.sum ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

# Собираем сервис
WORKDIR /app/auth/cmd/auth
RUN go build -o auth_service

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/auth/cmd/auth/auth_service .
COPY --from=builder /app/auth/config/config.yaml ./config/config.yaml
COPY --from=builder /app/auth/db/migrations /root/db/migrations

EXPOSE 8080

CMD goose -dir /root/db/migrations postgres "postgresql://${APP_POSTGRES_USER}:${APP_POSTGRES_PASSWORD}@${APP_POSTGRES_ADDRESS}/${APP_POSTGRES_NAME}?sslmode=disable" up && ./auth_service

