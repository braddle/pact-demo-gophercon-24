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

	err := pact.VerifyProvider(t, provider.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost%s", os.Getenv("HOST")),
		BrokerURL:                  "https://testingallthethings.pactflow.io/",
		Provider:                   "ScoopDash",
		BrokerToken:                "HwK7pR4-GTVCCPjN8a_JWw",
		ProviderVersion:            "1.0",
		PublishVerificationResults: publish,
	})

	assert.NoError(t, err)
}
