.PHONY: build run test

build:
	docker-compose build

up:
	docker-compose up -d

bash:
	docker-compose exec app bin/sh

test:
	go test ./tests/...
