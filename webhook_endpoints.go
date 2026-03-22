package meshpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// WebhookEndpointsResource handles webhook endpoint operations.
type WebhookEndpointsResource struct{ client *Client }

func normalizeWebhookList(b []byte) ([]map[string]interface{}, error) {
	if len(b) == 0 {
		return []map[string]interface{}{}, nil
	}
	var arr []map[string]interface{}
	if err := json.Unmarshal(b, &arr); err == nil {
		return arr, nil
	}
	var wrap struct {
		Data []map[string]interface{} `json:"data"`
	}
	if err := json.Unmarshal(b, &wrap); err != nil {
		return nil, err
	}
	if wrap.Data == nil {
		return []map[string]interface{}{}, nil
	}
	return wrap.Data, nil
}

// List returns the list of webhook endpoints.
func (r *WebhookEndpointsResource) List() ([]map[string]interface{}, error) {
	b, err := r.client.do("GET", "/webhook-endpoints", nil, "")
	if err != nil {
		return nil, err
	}
	return normalizeWebhookList(b)
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
func (r *WebhookEndpointsResource) Create(u string, events []string, secret string) (map[string]interface{}, error) {
	body := map[string]interface{}{"url": u, "events": events}
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

// Update updates a webhook endpoint (active, events per OpenAPI).
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

// ListDeliveries lists webhook deliveries for an endpoint.
func (r *WebhookEndpointsResource) ListDeliveries(id string, limit *int) ([]map[string]interface{}, error) {
	path := fmt.Sprintf("/webhook-endpoints/%s/deliveries", id)
	if limit != nil {
		path = path + "?limit=" + url.QueryEscape(fmt.Sprint(*limit))
	}
	b, err := r.client.do("GET", path, nil, "")
	if err != nil {
		return nil, err
	}
	return normalizeWebhookList(b)
}
