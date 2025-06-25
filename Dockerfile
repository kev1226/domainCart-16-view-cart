# Etapa 1: Build del binario
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compilación estática para Alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o get-cart .

# Etapa 2: Imagen final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/get-cart .

CMD ["./get-cart"]
