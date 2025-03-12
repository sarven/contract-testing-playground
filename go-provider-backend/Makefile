.PHONY: build up bash test migrate test-contract

build:
	docker compose build

up:
	docker compose up -d

bash:
	docker compose exec app /bin/bash

test:
	docker compose exec app go test ./tests/integration/...

test-contract:
	docker compose exec -e PACT_BROKER_URL=$$PACT_BROKER_URL -e PACT_BROKER_TOKEN=$$PACT_BROKER_TOKEN -e PROVIDER_VERSION=$$PROVIDER_VERSION app go test ./tests/pact/...

migrate:
	docker compose exec app sh -c 'migrate -database "$$DATABASE_URL" -source file://migrations up'
