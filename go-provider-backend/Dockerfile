FROM golang:1.24-alpine

RUN apk add --no-cache bash curl postgresql-client

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go
RUN ls -la

EXPOSE 8080

CMD ["./main"]
