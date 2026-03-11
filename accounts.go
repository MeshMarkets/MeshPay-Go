package meshpay

import "encoding/json"

// AccountsResource handles account operations.
type AccountsResource struct{ client *Client }

// Create creates an end-user account.
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
