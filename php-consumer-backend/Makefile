build:
	docker compose build

composer-install:
	docker compose exec app composer install

up:
	docker compose up -d

bash:
	docker compose exec app /bin/sh

run-migrations:
	docker compose exec app bin/console d:s:u --force --env=test

test-contract:
	docker compose exec app vendor/bin/phpunit --config=phpunit.xml.dist --testsuite "Contract Tests"
