package meshpay

import "encoding/json"

// HealthResource handles health check.
type HealthResource struct{ client *Client }

// Get returns the health status.
func (r *HealthResource) Get() (map[string]interface{}, error) {
	b, err := r.client.doNoAuth("GET", "/health")
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return map[string]interface{}{}, nil
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
