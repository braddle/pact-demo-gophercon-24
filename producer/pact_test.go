package producer_test

import (
	"fmt"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/pact-foundation/pact-go/v2/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"net"
	"net/http"
	"os"
	"scoop-dash/producer"
	"testing"
)

var port, _ = utils.GetFreePort()

func TestPacts(t *testing.T) {
	pact := provider.NewVerifier()

	go startTestServer()

	publish := os.Getenv("PUBLISH") == "yes"

	token, _ := os.ReadFile("/run/secrets/pactflow_token")

	err := pact.VerifyProvider(t, provider.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		BrokerURL:                  "https://testingallthethings.pactflow.io/",
		Provider:                   "ScoopDash",
		BrokerToken:                string(token),
		ProviderVersion:            "1.1",
		PublishVerificationResults: publish,
	})

	assert.NoError(t, err)
}

func startTestServer() {
	mux := producer.GetRouter()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("API starting: port %d (%s)", port, ln.Addr())
	log.Printf("API terminating: %v", http.Serve(ln, mux))
}
