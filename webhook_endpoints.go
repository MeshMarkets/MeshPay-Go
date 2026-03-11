package meshpay

import (
	"encoding/json"
	"fmt"
)

// WebhookEndpointsResource handles webhook endpoint operations.
type WebhookEndpointsResource struct{ client *Client }

// List returns the list of webhook endpoints.
func (r *WebhookEndpointsResource) List() ([]map[string]interface{}, error) {
	b, err := r.client.do("GET", "/webhook-endpoints", nil, "")
	if err != nil {
		return nil, err
	}
	var out []map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Get returns a single webhook endpoint.
func (r *WebhookEndpointsResource) Get(id string) (map[string]interface{}, error) {
	b, err := r.client.do("GET", fmt.Sprintf("/webhook-endpoints/%s", id), nil, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a webhook endpoint.
func (r *WebhookEndpointsResource) Create(url string, events []string, secret string) (map[string]interface{}, error) {
	body := map[string]interface{}{"url": url, "events": events}
	if secret != "" {
		body["secret"] = secret
	}
	b, err := r.client.do("POST", "/webhook-endpoints", body, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update updates a webhook endpoint (active, events).
func (r *WebhookEndpointsResource) Update(id string, active *bool, events []string) (map[string]interface{}, error) {
	body := map[string]interface{}{}
	if active != nil {
		body["active"] = *active
	}
	if events != nil {
		body["events"] = events
	}
	b, err := r.client.do("PATCH", fmt.Sprintf("/webhook-endpoints/%s", id), body, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes a webhook endpoint.
func (r *WebhookEndpointsResource) Delete(id string) error {
	_, err := r.client.do("DELETE", fmt.Sprintf("/webhook-endpoints/%s", id), nil, "")
	return err
}
