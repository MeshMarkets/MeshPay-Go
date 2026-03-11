package meshpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// PayoutsResource handles payout operations.
type PayoutsResource struct{ client *Client }

// List returns a paginated list of payouts.
func (r *PayoutsResource) List(opts *ListOptions) (map[string]interface{}, error) {
	path := "/payouts"
	if opts != nil {
		v := url.Values{}
		if opts.Limit != nil {
			v.Set("limit", fmt.Sprint(*opts.Limit))
		}
		if opts.Cursor != nil {
			v.Set("cursor", *opts.Cursor)
		}
		if opts.Status != nil {
			v.Set("status", *opts.Status)
		}
		if len(v) > 0 {
			path = path + "?" + v.Encode()
		}
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

// Get returns a single payout.
func (r *PayoutsResource) Get(payoutID string) (map[string]interface{}, error) {
	b, err := r.client.do("GET", fmt.Sprintf("/payouts/%s", payoutID), nil, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a payout.
func (r *PayoutsResource) Create(accountID string, amount float64, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/payouts", map[string]interface{}{
		"account_id": accountID,
		"amount":     amount,
	}, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
