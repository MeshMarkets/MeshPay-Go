package meshpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const defaultBaseURL = "https://YOUR_PROJECT_REF.supabase.co/functions/v1/api"

// Client for the Mesh Pay API.
type Client struct {
	baseURL      string
	apiKey       string
	client       *http.Client
	useXApiKey   bool
}

// New creates a new Client.
func New(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		apiKey:  apiKey,
		client:  &http.Client{},
	}
}

// NewWithOptions creates a client with optional X-Api-Key header instead of Bearer.
func NewWithOptions(apiKey, baseURL string, useXApiKey bool) *Client {
	c := New(apiKey, baseURL)
	c.useXApiKey = useXApiKey
	return c
}

func (c *Client) do(method, path string, body interface{}, idempotencyKey string) ([]byte, error) {
	return c.doInternal(method, path, body, idempotencyKey, false)
}

func (c *Client) doNoAuth(method, path string) ([]byte, error) {
	return c.doInternal(method, path, nil, "", true)
}

func (c *Client) doInternal(method, path string, body interface{}, idempotencyKey string, skipAuth bool) ([]byte, error) {
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
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if !skipAuth {
		if c.useXApiKey {
			req.Header.Set("X-Api-Key", c.apiKey)
		} else {
			req.Header.Set("Authorization", "Bearer "+c.apiKey)
		}
	}
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
	if resp.StatusCode == 204 || len(b) == 0 {
		return nil, nil
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
