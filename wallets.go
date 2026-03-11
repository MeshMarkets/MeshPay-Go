package meshpay

import (
	"encoding/json"
	"fmt"
)

// WalletsResource handles wallet operations.
type WalletsResource struct{ client *Client }

// Create gets or creates a wallet for the account.
func (r *WalletsResource) Create(accountID string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/wallets", map[string]string{"account_id": accountID}, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetByAccountID returns the wallet for the account.
func (r *WalletsResource) GetByAccountID(accountID string) (map[string]interface{}, error) {
	b, err := r.client.do("GET", fmt.Sprintf("/wallets?account_id=%s", accountID), nil, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
