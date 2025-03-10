test-unit:
	docker compose exec frontend npm run test:unit

test-contract:
	docker compose exec frontend npm run test:contract

publish-contracts:
	docker compose exec frontend pact-broker publish /app/pacts/ReactFrontend-Backend.json --consumer-app-version=$$CONSUMER_VERSION --broker-base-url=$$PACT_BROKER_URL --broker-token=$$PACT_BROKER_TOKEN

