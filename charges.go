package meshpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ChargesResource handles charge operations.
type ChargesResource struct{ client *Client }

// List returns a paginated list of charges.
func (r *ChargesResource) List(opts *ListOptions) (map[string]interface{}, error) {
	path := "/charges"
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

// Get returns a single charge.
func (r *ChargesResource) Get(chargeID string) (map[string]interface{}, error) {
	b, err := r.client.do("GET", fmt.Sprintf("/charges/%s", chargeID), nil, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates a charge.
func (r *ChargesResource) Create(body map[string]interface{}, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/charges", body, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Fund funds a charge.
func (r *ChargesResource) Fund(chargeID string, body map[string]interface{}, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/charges/%s/fund", chargeID), body, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Cancel cancels a pending charge.
func (r *ChargesResource) Cancel(chargeID, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/charges/%s/cancel", chargeID), map[string]interface{}{}, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Refund refunds a charge.
func (r *ChargesResource) Refund(chargeID, idempotencyKey string, body map[string]interface{}) (map[string]interface{}, error) {
	if body == nil {
		body = map[string]interface{}{}
	}
	b, err := r.client.do("POST", fmt.Sprintf("/charges/%s/refund", chargeID), body, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
