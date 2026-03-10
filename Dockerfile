FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

ENV GONOSUMDB="github.com/InBitGT/*"
ENV GOPRIVATE="github.com/InBitGT/*"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /root/

# Copiar el binario 
COPY --from=builder /app/app .

EXPOSE 8086
CMD ["./app"]