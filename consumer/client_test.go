package consumer_test

import (
	"fmt"
	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
		Given("There is an ice cream white-chocolate-magnum").
		UponReceiving("A request for an ice cream with ID white-chocolate-magnum").
		WithRequestPathMatcher(http.MethodGet, matchers.Regex("/icecream/white-chocolate-magnum", "\\/icecream\\/[a-z0-9-]+"), func(b *consumer.V2RequestBuilder) {
			b.Header("Accept", matchers.Equality("application/json"))
		}).
		//WithRequest(http.MethodGet, "/icecream/white-chocolate-magnum", func(b *consumer.V2RequestBuilder) {
		//	b.Header("Accept", matchers.Equality("application/json"))
		//}).
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
		//assert.Equal(t, int64(2020), ic.YearOfRelease)

		return err
	})

	assert.NoError(t, err)
}

func TestHandleInvalidHost(t *testing.T) {
	c := client.NewClient("::::::::")

	_, err := c.GetIceCream("white-chocolate-magnum")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Error creating request")
}

func TestHandleConnectionFailure(t *testing.T) {
	c := client.NewClient("http://localhost:1111")

	_, err := c.GetIceCream("white-chocolate-magnum")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Error making request to server")
}

func TestHandlesInvalidContentInResponse(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<json>INVALID</json>"))
	}))
	defer s.Close()

	c := client.NewClient(s.URL)

	_, err := c.GetIceCream("white-chocolate-magnum")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid response body")
}

func TestGetIceCreamThatDoesNotExists(t *testing.T) {
	pact.AddInteraction().
		UponReceiving("A request for an ice cream with ID coffee-and-chocolate-magnum").
		WithRequest(http.MethodGet, "/icecream/coffee-and-chocolate-magnum", func(b *consumer.V2RequestBuilder) {
			b.Header("Accept", matchers.Equality("application/json"))
		}).
		WillRespondWith(http.StatusNotFound, func(b *consumer.V2ResponseBuilder) {
			b.BodyMatch(&client.HTTPError{})
		})

	err := pact.ExecuteTest(t, func(config consumer.MockServerConfig) error {
		c := client.NewClient(fmt.Sprintf("http://%s:%d", config.Host, config.Port))

		_, err := c.GetIceCream("coffee-and-chocolate-magnum")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "No ice cream with ID: coffee-and-chocolate-magnum (404)")

		return nil
	})

	assert.NoError(t, err)
}

func TestMain(m *testing.M) {
	beforeAllTests()
	code := m.Run()
	os.Exit(code)
}
