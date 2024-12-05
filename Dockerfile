FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN go build -o bin/akaSocial main.go

FROM debian:bullseye-slim

ENV APP_BINARY=/app/bin/akaSocial \
    APP_PORT=8080

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/bin/akaSocial ./bin/

EXPOSE 8080

CMD ["./bin/akaSocial"]
