test-unit:
	docker compose run --rm frontend npm run test:unit

test-contract:
	docker compose run --rm frontend npm run test:contract

publish-contracts:
	docker compose run --rm -e PACT_BROKER_URL -e PACT_BROKER_TOKEN -e CONSUMER_VERSION frontend pact-broker publish pacts --consumer-app-version $$CONSUMER_VERSION --broker-base-url $$PACT_BROKER_URL --broker-token $$PACT_BROKER_TOKEN
