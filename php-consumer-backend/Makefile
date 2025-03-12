build:
	docker compose build

composer-install:
	docker compose exec app composer install

up:
	docker compose up -d

bash:
	docker compose exec app /bin/sh

test-contract:
	docker compose exec app vendor/bin/phpunit --config=phpunit.xml.dist --testsuite "Contract Tests"
