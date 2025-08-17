FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y default-mysql-client \
    && wget https://github.com/golang-migrate/migrate/releases/download/v4.16.0/migrate.linux-amd64.tar.gz \
    && tar -xvzf migrate.linux-amd64.tar.gz \
    && mv migrate /usr/local/bin/migrate \
    && rm migrate.linux-amd64.tar.gz \
    && chmod +x /usr/local/bin/migrate \
    && rm -rf /var/lib/apt/lists/*

RUN go build -o main ./cmd/main.go

RUN chmod +x ./docker-entrypoint.sh

# CMD ["./main"]
ENTRYPOINT ["./docker-entrypoint.sh"]
