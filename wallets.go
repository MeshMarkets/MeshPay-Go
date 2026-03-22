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

// ListFiatAccounts lists linked fiat accounts for a membership.
func (r *WalletsResource) ListFiatAccounts(membershipID string) (map[string]interface{}, error) {
	path := "/wallets/fiat-accounts?membership_id=" + url.QueryEscape(membershipID)
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

// LinkFiatAccount links a fiat account (requires Idempotency-Key).
func (r *WalletsResource) LinkFiatAccount(body map[string]interface{}, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/wallets/fiat-accounts", body, idempotencyKey)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return map[string]interface{}{}, nil
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// UnlinkFiatAccount removes a linked fiat account.
func (r *WalletsResource) UnlinkFiatAccount(membershipID, fiatAccountID, idempotencyKey string) error {
	path := fmt.Sprintf(
		"/wallets/fiat-accounts?membership_id=%s&fiat_account_id=%s",
		url.QueryEscape(membershipID),
		url.QueryEscape(fiatAccountID),
	)
	_, err := r.client.do("DELETE", path, nil, idempotencyKey)
	return err
}
