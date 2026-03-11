package meshpay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// VerifySignature verifies the HMAC-SHA256 signature of a webhook payload.
func VerifySignature(payload, signature, secret string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(signature))
}
