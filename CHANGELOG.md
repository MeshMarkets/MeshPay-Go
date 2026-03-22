# Changelog

## 2.0.0

Breaking — OpenAPI v1 alignment. Removed `Payouts` and `APIKeys`. Default base URL: `https://YOUR_PROJECT_REF.supabase.co/functions/v1/api`. Use `NewWithOptions(apiKey, baseURL, useXApiKey bool)` for `X-Api-Key` auth. Escrows list uses `EscrowListOptions` (limit, status). Charges `Create` takes a body map; `Fund`/`Refund`/`Cancel` updated; `OpenDispute`/`ResolveDispute` added. On-ramp/Off-ramp: `CreateSession`.
