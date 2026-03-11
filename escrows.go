package meshpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// EscrowsResource handles escrow operations.
type EscrowsResource struct{ client *Client }

// List returns a paginated list of escrows.
func (r *EscrowsResource) List(opts *ListOptions) (map[string]interface{}, error) {
	path := "/escrows"
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

// Get returns a single escrow.
func (r *EscrowsResource) Get(escrowID string) (map[string]interface{}, error) {
	b, err := r.client.do("GET", fmt.Sprintf("/escrows/%s", escrowID), nil, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Release releases the escrow to the seller.
func (r *EscrowsResource) Release(escrowID, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/escrows/%s/release", escrowID), map[string]interface{}{}, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
