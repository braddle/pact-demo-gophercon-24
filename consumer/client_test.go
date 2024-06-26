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

func TestGetIceCreamWhiteChocolateMagnum(t *testing.T) {
	pact.AddInteraction().
		UponReceiving("A request for an ice cream that does not exist").
		WithRequest(http.MethodGet, "/icecream/white-chocolate-magnum", func(b *consumer.V2RequestBuilder) {
			b.Header("Accept", matchers.Equality("application/json"))
		}).
		WillRespondWith(http.StatusOK, func(b *consumer.V2ResponseBuilder) {
			b.BodyMatch(&client.IceCream{})
		})

	err := pact.ExecuteTest(t, func(config consumer.MockServerConfig) error {
		c := client.NewClient(fmt.Sprintf("http://%s:%d", config.Host, config.Port))

		ic, err := c.GetIceCream("white-chocolate-magnum")

		assert.NoError(t, err)
		assert.Equal(t, "white-chocolate-magnum", ic.ID)
		assert.Equal(t, "1234567891234", ic.Barcode)
		assert.Equal(t, "White Chocolate Magnum", ic.Name)
		assert.Equal(t, "Walls", ic.Manufacturer.Name)
		assert.Equal(t, "Unilever UK, ShareHappy, Freepost ADM3940, London, SW1A 1YR", ic.Manufacturer.Address)
		assert.Equal(t, []string{"string", "string"}, ic.Ingredients)
		assert.Equal(t, int64(227), ic.Calories)
		assert.True(t, ic.RecyclablePackaging)
		assert.Equal(t, float64(4.2), ic.Rating)
		assert.Len(t, ic.Images, 1)
		assert.Equal(t, "https://www.eden-farm.co.uk/media/catalog/product/cache/e70602422d911f0edb0b0d50a9ac95bc/f/1/f1300eec6c4489d61d37f2c6b91602e1.jpg", ic.Images[0].URL)
		assert.Equal(t, int64(700), ic.Images[0].Width)
		assert.Equal(t, int64(700), ic.Images[0].Height)

		return err
	})

	assert.NoError(t, err)
}

func TestMain(m *testing.M) {
	beforeAllTests()
	code := m.Run()
	os.Exit(code)
}
