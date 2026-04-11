package meshpay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// EscrowListOptions lists escrows (limit, status only per OpenAPI).
type EscrowListOptions struct {
	Limit  *int
	Status *string
}

// EscrowsResource handles escrow operations.
type EscrowsResource struct{ client *Client }

// List returns a paginated list of escrows.
func (r *EscrowsResource) List(opts *EscrowListOptions) (map[string]interface{}, error) {
	path := "/escrows"
	if opts != nil {
		v := url.Values{}
		if opts.Limit != nil {
			v.Set("limit", fmt.Sprint(*opts.Limit))
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

// CreateContribution adds a contribution to a pooled escrow.
func (r *EscrowsResource) CreateContribution(escrowID string, body map[string]interface{}, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/escrows/%s/contributions", escrowID), body, idempotencyKey)
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

// SetPayee assigns the payee for a pooled escrow.
func (r *EscrowsResource) SetPayee(escrowID string, body map[string]interface{}, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/escrows/%s/set-payee", escrowID), body, idempotencyKey)
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

// CancelPooledEscrow cancels a pending pooled escrow and refunds contributors.
func (r *EscrowsResource) CancelPooledEscrow(escrowID, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/escrows/%s/cancel-pool", escrowID), map[string]interface{}{}, idempotencyKey)
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

// OpenDispute records an on-chain dispute.
func (r *EscrowsResource) OpenDispute(escrowID, txHash string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/escrows/%s/open-dispute", escrowID), map[string]string{"tx_hash": txHash}, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ResolveDispute resolves a dispute on-chain.
func (r *EscrowsResource) ResolveDispute(escrowID string, releaseToSeller bool, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", fmt.Sprintf("/escrows/%s/resolve-dispute", escrowID), map[string]bool{"release_to_seller": releaseToSeller}, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
