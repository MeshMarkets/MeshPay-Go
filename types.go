package meshpay

// ListOptions is used for list endpoints (limit, cursor, status).
type ListOptions struct {
	Limit  *int
	Cursor *string
	Status *string
}
