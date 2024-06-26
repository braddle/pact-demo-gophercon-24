package main_test

import (
	"fmt"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPacts(t *testing.T) {
	pact := provider.NewVerifier()

	publish := os.Getenv("PUBLISH") == "yes"

	token, _ := os.ReadFile("/run/secrets/pactflow_token")

	err := pact.VerifyProvider(t, provider.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost%s", os.Getenv("HOST")),
		BrokerURL:                  "https://testingallthethings.pactflow.io/",
		Provider:                   "ScoopDash",
		BrokerToken:                string(token),
		ProviderVersion:            "1.1",
		PublishVerificationResults: publish,
	})

	assert.NoError(t, err)
}
