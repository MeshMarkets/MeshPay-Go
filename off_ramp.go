package meshpay

import (
	"encoding/json"
)

// OffRampResource handles off-ramp (USDC → USD) operations.
type OffRampResource struct{ client *Client }

// GetQuote returns a quote for the given amount.
func (r *OffRampResource) GetQuote(amountUSDC, amountUSD *float64) (map[string]interface{}, error) {
	body := map[string]interface{}{}
	if amountUSDC != nil {
		body["amount_usdc"] = *amountUSDC
	}
	if amountUSD != nil {
		body["amount_usd"] = *amountUSD
	}
	b, err := r.client.do("POST", "/off-ramp/quote", body, "")
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ExecuteTrade executes an off-ramp trade with the given quote ID.
func (r *OffRampResource) ExecuteTrade(quoteID, idempotencyKey string) (map[string]interface{}, error) {
	b, err := r.client.do("POST", "/off-ramp/trade", map[string]string{"quote_id": quoteID}, idempotencyKey)
	if err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}
