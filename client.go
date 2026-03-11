package meshpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client for the Mesh Pay API.
type Client struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

// New creates a new Client. Pass empty baseURL to use http://localhost:3000.
func New(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = "http://localhost:3000"
	}
	return &Client{baseURL: baseURL, apiKey: apiKey, client: &http.Client{}}
}

// do performs an HTTP request and returns the body. Returns an error on 4xx/5xx.
func (c *Client) do(method, path string, body interface{}, idempotencyKey string) ([]byte, error) {
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		r = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, c.baseURL+path, r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	if idempotencyKey != "" {
		req.Header.Set("Idempotency-Key", idempotencyKey)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("meshpay api error %d: %s", resp.StatusCode, string(b))
	}
	return b, nil
}

// Health returns the health resource.
func (c *Client) Health() *HealthResource {
	return &HealthResource{client: c}
}

// Accounts returns the accounts resource.
func (c *Client) Accounts() *AccountsResource {
	return &AccountsResource{client: c}
}

// Wallets returns the wallets resource.
func (c *Client) Wallets() *WalletsResource {
	return &WalletsResource{client: c}
}

// Charges returns the charges resource.
func (c *Client) Charges() *ChargesResource {
	return &ChargesResource{client: c}
}

// Escrows returns the escrows resource.
func (c *Client) Escrows() *EscrowsResource {
	return &EscrowsResource{client: c}
}

// Payouts returns the payouts resource.
func (c *Client) Payouts() *PayoutsResource {
	return &PayoutsResource{client: c}
}

// APIKeys returns the API keys resource.
func (c *Client) APIKeys() *APIKeysResource {
	return &APIKeysResource{client: c}
}

// WebhookEndpoints returns the webhook endpoints resource.
func (c *Client) WebhookEndpoints() *WebhookEndpointsResource {
	return &WebhookEndpointsResource{client: c}
}

// OnRamp returns the on-ramp resource.
func (c *Client) OnRamp() *OnRampResource {
	return &OnRampResource{client: c}
}

// OffRamp returns the off-ramp resource.
func (c *Client) OffRamp() *OffRampResource {
	return &OffRampResource{client: c}
}
