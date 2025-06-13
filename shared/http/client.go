package httphelper

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shared/config"
	"shared/constant"
	shared_context "shared/context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type ResponseWrapper[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func DecodeResponse[T any](resp *http.Response) (*ResponseWrapper[T], error) {
	defer resp.Body.Close()
	var wrapper ResponseWrapper[T]
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, err
	}
	return &wrapper, nil
}

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

func (c *Client) DoRaw(ctx context.Context, serviceName, method, path string, in any) (*http.Response, error) {
	baseURL, ok := c.services[serviceName]
	if !ok {
		return nil, errors.New("unregistered service: " + serviceName)
	}

	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, baseURL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	token, _ := shared_context.GetToken(ctx)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	if ip, ok := ctx.Value(constant.ContextIPAddress).(string); ok && ip != "" {
		req.Header.Set("X-Forwarded-For", ip)
	}
	if requestID, ok := ctx.Value(constant.ContextRequestID).(string); ok && requestID != "" {
		req.Header.Set("X-Request-ID", requestID)
	}
	if traceID, ok := ctx.Value(constant.ContextTraceID).(string); ok && traceID != "" {
		req.Header.Set("X-Trace-ID", traceID)
	}

	// Inject the OTel context into the outgoing HTTP headers
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		errMsg := "[" + serviceName + "] " + resp.Status
		resp.Body.Close()
		return nil, errors.New(errMsg)
	}

	return resp, nil
}

func DoAndDecode[T any](c *Client, ctx context.Context, serviceName, method, path string, in any) (*ResponseWrapper[T], error) {
	resp, err := c.DoRaw(ctx, serviceName, method, path, in)
	if err != nil {
		return nil, err
	}
	return DecodeResponse[T](resp)
}
