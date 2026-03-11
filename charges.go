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
func (r *ChargesResource) Create(buyerID, sellerAccountID string, amount float64, idempotencyKey string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"buyer_id":            buyerID,
		"seller_account_id":  sellerAccountID,
		"amount":              amount,
	}
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

// Fund funds a charge (entity_secret_ciphertext required).
func (r *ChargesResource) Fund(chargeID, entitySecretCiphertext, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/charges/%s/fund", chargeID),
		map[string]string{"entity_secret_ciphertext": entitySecretCiphertext}, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Refund refunds a charge (amount optional for partial).
func (r *ChargesResource) Refund(chargeID, idempotencyKey string, amount *float64) (map[string]interface{}, error) {
	body := map[string]interface{}{}
	if amount != nil {
		body["amount"] = *amount
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
