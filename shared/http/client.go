package httphelper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shared/config"
	"time"
)

type Client struct {
	services   map[string]string
	httpClient *http.Client
}

func NewWithServices() *Client {
	return &Client{
		services: GetServiceMap(RegisteredServers(config.LoadConfig())),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Do(ctx context.Context, serviceName, method, path string, in, out any) error {
	baseURL, ok := c.services[serviceName]
	if !ok {
		return errors.New("Unregistered service: " + serviceName)
	}

	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, baseURL+path, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.New("[" + serviceName + "] " + resp.Status)
	}

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}

	return nil
}
