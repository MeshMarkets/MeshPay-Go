package meshpay

import (
	"encoding/json"
)

// OnRampResource handles on-ramp (USD → USDC) operations.
type OnRampResource struct{ client *Client }

// GetQuote returns a quote for the given amount.
func (r *OnRampResource) GetQuote(amountUSD, amountUSDC *float64) (map[string]interface{}, error) {
	body := map[string]interface{}{}
	if amountUSD != nil {
		body["amount_usd"] = *amountUSD
	}
	if amountUSDC != nil {
		body["amount_usdc"] = *amountUSDC
	}
	b, err := r.client.do("POST", "/on-ramp/quote", body, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ExecuteTrade executes an on-ramp trade with the given quote ID.
func (r *OnRampResource) ExecuteTrade(quoteID, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/on-ramp/trade", map[string]string{"quote_id": quoteID}, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
