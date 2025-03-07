.PHONY: build up bash test migrate

build:
	docker compose build

up:
	docker compose up -d

bash:
	docker compose exec app /bin/sh

test:
	docker compose exec app go test ./tests/...

wait-for-app:
	echo "Waiting for app service to be healthy..."
	docker compose exec app sh -c 'while ! nc -z localhost 8080; do sleep 1; done'

migrate: wait-for-app
	docker compose exec app sh -c 'migrate -database "$$DATABASE_URL" -source file://migrations up'
