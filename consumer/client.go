package consumer

import (
	"fmt"
	"net/http"
)

type Client struct {
	host string
}

type HealthStatus struct {
	status string
}

func (c Client) healthcheck() (HealthStatus, error) {
	url := fmt.Sprintf("http://%s/health", c.host)
	fmt.Println(url)
	http.Get(url)
	return HealthStatus{status: "OK"}, nil
}

func NewClient(host string) Client {
	return Client{host: host}
}
