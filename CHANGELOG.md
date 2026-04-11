# Changelog

## 3.1.0

- **Added** `ChargesResource.CreatePooledCharge` and `EscrowsResource.CreateContribution`, `SetPayee`, `CancelPooledEscrow` for pooled charge and escrow routes.

## 3.0.0

- **Removed** `ListFiatAccounts`, `LinkFiatAccount`, `UnlinkFiatAccount` from wallets. Use `OnRamp` / `OffRamp` `CreateSession` instead.

## 2.0.0

Breaking — OpenAPI v1 alignment. Removed `Payouts` and `APIKeys`. Default base URL: `https://YOUR_PROJECT_REF.supabase.co/functions/v1/api`. Use `NewWithOptions(apiKey, baseURL, useXApiKey bool)` for `X-Api-Key` auth. Escrows list uses `EscrowListOptions` (limit, status). Charges `Create` takes a body map; `Fund`/`Refund`/`Cancel` updated; `OpenDispute`/`ResolveDispute` added. On-ramp/Off-ramp: `CreateSession`.
