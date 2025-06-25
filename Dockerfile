# Etapa 1: Compilar el binario
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# CAMBIA ESTO: Compila usando el nombre del servicio (por ejemplo: get-cart)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o get-cart .

# Etapa 2: Imagen final liviana
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# CAMBIA ESTO TAMBIÉN
COPY --from=builder /app/get-cart .

EXPOSE 3036

# CAMBIA ESTO IGUAL
CMD ["./get-cart"]
