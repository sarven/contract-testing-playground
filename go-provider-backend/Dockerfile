FROM golang:1.24-alpine

RUN apk add --no-cache bash curl postgresql-client musl musl-dev libffi-dev libffi-dev ruby ruby-dev build-base && \
gem install pact-cli

WORKDIR /app

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/pact-foundation/pact-go/v2@latest && \
    ln -s /go/bin/pact-go /usr/local/bin/pact-go

RUN pact-go -l DEBUG install

RUN export PATH=$PATH:/root/.pact/bin && \
    echo 'export PATH=$PATH:/root/.pact/bin' >> ~/.bashrc

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
