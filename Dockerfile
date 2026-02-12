FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /root/

# Copiar el binario
COPY --from=builder /app/app .

# Copiar el archivo .env
COPY .env .

EXPOSE 8086
CMD ["./app"]