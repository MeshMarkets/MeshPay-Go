package meshpay

import (
	"encoding/json"
	"fmt"
)

// APIKeysResource handles API key operations.
type APIKeysResource struct{ client *Client }

// List returns the list of API keys.
func (r *APIKeysResource) List() ([]map[string]interface{}, error) {
	b, err := r.client.do("GET", "/api-keys", nil, "")
	if err != nil {
		return nil, err
	}
	var out []map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates an API key. Optional name.
func (r *APIKeysResource) Create(name string) (map[string]interface{}, error) {
	body := map[string]interface{}{}
	if name != "" {
		body["name"] = name
	}
	b, err := r.client.do("POST", "/api-keys", body, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete deletes an API key.
func (r *APIKeysResource) Delete(id string) error {
	_, err := r.client.do("DELETE", fmt.Sprintf("/api-keys/%s", id), nil, "")
	return err
}
