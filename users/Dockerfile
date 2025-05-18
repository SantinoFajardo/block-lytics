# Etapa de build
FROM golang:1.23.6-alpine AS builder


WORKDIR /app
COPY users/ .
COPY proto ../proto


RUN go mod tidy && \
    go build -o app ./cmd/main.go

# Imagen final
FROM alpine:latest

WORKDIR /root/
RUN apk add --no-cache postgresql-client
COPY --from=builder /app/app .
COPY users/wait-for.sh /wait-for.sh

EXPOSE 50052
ENTRYPOINT ["/wait-for.sh", "postgres", "./app"]
