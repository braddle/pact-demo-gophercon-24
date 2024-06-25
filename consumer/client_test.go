package consumer_test

import (
	"fmt"
	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	client "scoop-dash/consumer"
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
		WithRequest(http.MethodGet, "/health", func(b *consumer.V2RequestBuilder) {
			b.Header("Accept", matchers.Equality("application/json"))
		}).
		WillRespondWith(http.StatusOK, func(b *consumer.V2ResponseBuilder) {
			b.BodyMatch(&client.HealthStatus{})
		})

	err := pact.ExecuteTest(t, func(config consumer.MockServerConfig) error {

		c := client.NewClient(fmt.Sprintf("http://%s:%d", config.Host, config.Port))
		health, err := c.Healthcheck()

		assert.NoError(t, err)
		assert.Equal(t, "OK", health.Status)

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
