package meshpay

import "encoding/json"

// OnRampResource handles on-ramp (USD → USDC) operations.
type OnRampResource struct{ client *Client }

// CreateSession POST /on-ramp/sessions (Coinbase-hosted onramp).
func (r *OnRampResource) CreateSession(body map[string]interface{}) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/on-ramp/sessions", body, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
