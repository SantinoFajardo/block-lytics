# Etapa de build
FROM golang:1.23.6-alpine AS builder

WORKDIR /app
COPY gateway/ . 
COPY proto ../proto         

RUN go mod tidy && \
    go build -o app ./cmd/main.go

# Imagen final
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
