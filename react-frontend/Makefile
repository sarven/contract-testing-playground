test-unit:
	docker compose run --rm frontend npm run test:unit

test-contract:
	docker compose run --rm frontend npm run test:contract

publish-contracts:
	pact-broker publish pacts --consumer-app-version $(CONSUMER_VERSION) --broker-base-url $(PACT_BROKER_URL) --broker-token $(PACT_BROKER_TOKEN)
