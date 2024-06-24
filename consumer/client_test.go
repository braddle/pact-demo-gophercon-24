package consumer

import (
	"fmt"
	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var pact *consumer.V2HTTPMockProvider

func beforeAllTests() {
	pact, _ = consumer.NewV2Pact(consumer.MockHTTPProviderConfig{
		Consumer: "ScoopDashClient",
		Provider: "ScoopDash",
	})
}

func TestHealthcheck(t *testing.T) {
	pact.AddInteraction().
		Given("The SccopDash service is up an health").
		UponReceiving("A health check request").
		WithRequest(http.MethodGet, "/health").
		WillRespondWith(http.StatusOK)

	err := pact.ExecuteTest(t, func(config consumer.MockServerConfig) error {
		client := NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port))
		health, err := client.healthcheck()

		assert.NoError(t, err)
		assert.Equal(t, "OK", health.status)

		return err
	})

	assert.NoError(t, err)
}

func TestMain(m *testing.M) {
	beforeAllTests()
	code := m.Run()
	//afterAllTests()
	os.Exit(code)
}
