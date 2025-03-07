.PHONY: build up bash test migrate

build:
	docker compose build

up:
	docker compose up -d

bash:
	docker compose exec app /bin/sh

test:
	docker compose exec app go test ./tests/...

migrate:
	docker compose exec app sh -c 'migrate -database "$$DATABASE_URL" -source file://migrations up'
