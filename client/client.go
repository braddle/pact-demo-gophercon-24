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
	url := fmt.Sprintf("http://%s/health", c.host)

	resp, _ := http.Get(url)
	hs := HealthStatus{}
	json.NewDecoder(resp.Body).Decode(&hs)
	return hs, nil
}

func NewClient(host string) Client {
	return Client{host: host}
}
