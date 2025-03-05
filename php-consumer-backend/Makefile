build:
	docker compose build

up:
	docker compose up -d

bash:
	docker compose exec app /bin/sh

run-migrations:
	docker compose exec app bin/console d:s:u --force --env=test

test:
	docker compose exec app vendor/bin/phpunit

test-contract:
	docker compose exec app vendor/bin/phpunit --config=phpunit.xml.dist --testsuite "Contract Tests"
