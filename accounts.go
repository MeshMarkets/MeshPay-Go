package meshpay

import (
	"encoding/json"
	"fmt"
)

// AccountsResource handles account operations.
type AccountsResource struct{ client *Client }

// List lists memberships for the project.
func (r *AccountsResource) List() (map[string]interface{}, error) {
	b, err := r.client.do("GET", "/accounts", nil, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// Create creates an end-user account / membership.
func (r *AccountsResource) Create(email string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/accounts", map[string]string{"email": email}, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// DeleteMembership removes this project's membership.
func (r *AccountsResource) DeleteMembership(membershipID string) error {
	_, err := r.client.do("DELETE", fmt.Sprintf("/accounts/%s", membershipID), nil, "")
	return err
}
