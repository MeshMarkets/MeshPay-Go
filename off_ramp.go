package meshpay

import "encoding/json"

// OffRampResource handles off-ramp (USDC → USD) operations.
type OffRampResource struct{ client *Client }

// CreateSession POST /off-ramp/sessions (Coinbase-hosted offramp).
func (r *OffRampResource) CreateSession(body map[string]interface{}) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/off-ramp/sessions", body, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
