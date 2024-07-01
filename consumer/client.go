package consumer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	host string
}

type HealthStatus struct {
	Status string `json:"status" pact:"example=OK"`
}

type HTTPError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type IceCream struct {
	ID                  string       `json:"id" pact:"example=white-chocolate-magnum"`
	Barcode             string       `json:"barcode" pact:"example=1234567891234,regex=^(?:[0-9]{13}|[0-9]{8})$"`
	Name                string       `json:"name" pact:"example=White Chocolate Magnum"`
	Manufacturer        Manufacturer `json:"manufacturer"`
	Ingredients         []string     `json:"ingredients" pact:"min=2"`
	Calories            int64        `json:"calories" pact:"example=227"`
	RecyclablePackaging bool         `json:"recyclable_packaging" pact:"example=true"`
	Rating              float64      `json:"rating" pact:"example=4.2"`
	Images              []Image      `json:"images"`
}

type Image struct {
	URL    string `json:"url" pact:"example=https://www.eden-farm.co.uk/media/catalog/product/cache/e70602422d911f0edb0b0d50a9ac95bc/f/1/f1300eec6c4489d61d37f2c6b91602e1.jpg"`
	Width  int64  `json:"width" pact:"example=700"`
	Height int64  `json:"height" pact:"example=700"`
}

type Manufacturer struct {
	ID      string `json:"id" pact:"example=walls-unilever-uk"`
	Name    string `json:"name" pact:"example=Walls"`
	Address string `json:"address" pact:"example=Unilever UK, ShareHappy, Freepost ADM3940, London, SW1A 1YR"`
}

func (c Client) Healthcheck() (HealthStatus, error) {
	url := fmt.Sprintf("%s/health", c.host)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, _ := client.Do(req)

	hs := HealthStatus{}
	json.NewDecoder(resp.Body).Decode(&hs)
	return hs, nil
}

func (c Client) GetIceCream(id string) (IceCream, error) {
	ic := IceCream{}
	url := fmt.Sprintf("%s/icecream/white-chocolate-magnum", c.host)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ic, fmt.Errorf("Error creating request: %s", err.Error())
	}
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ic, fmt.Errorf("Error making request to server: %s", err.Error())
	}

	err = json.NewDecoder(resp.Body).Decode(&ic)
	if err != nil {
		return ic, fmt.Errorf("Invalid response body: %s", err.Error())
	}
	return ic, nil
}

func NewClient(host string) Client {
	return Client{host: host}
}
