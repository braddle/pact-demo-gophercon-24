test:
	go test -v ./...

publish:
	pact-broker publish ./pacts --consumer-app-version $$(uuid) --broker-base-url https://testingallthethings.pactflow.io/ --broker-token=$$(cat /run/secrets/pactflow_token)

can-i-deploy:
	pact-broker can-i-deploy --pacticipant ScoopDashClient --latest --to-environment production --broker-base-url https://testingallthethings.pactflow.io/ --broker-token=$$(cat /run/secrets/pactflow_token)