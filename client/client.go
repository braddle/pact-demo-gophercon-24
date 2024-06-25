package client

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

func NewClient(host string) Client {
	return Client{host: host}
}
