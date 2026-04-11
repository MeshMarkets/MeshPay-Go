package meshpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// WalletsResource handles wallet operations.
type WalletsResource struct{ client *Client }

// List returns wallet summaries for the project.
func (r *WalletsResource) List() (map[string]interface{}, error) {
	b, err := r.client.do("GET", "/wallets", nil, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetDetail returns wallet detail and token balances.
func (r *WalletsResource) GetDetail(membershipID string, network string) (map[string]interface{}, error) {
	path := fmt.Sprintf("/wallets/%s", membershipID)
	if network != "" {
		path = path + "?network=" + url.QueryEscape(network)
	}
	b, err := r.client.do("GET", path, nil, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
