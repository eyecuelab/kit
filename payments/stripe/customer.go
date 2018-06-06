package stripe

// Customer ...
type Customer struct {
	StripeCustomerID string
}

// Parent parent model interface
type Parent interface {
	StripeCustomerDescription() *string
	GetStripeCustomerID() *string
	SetStripeCustomerID(string) error
}

// GetStripeCustomerID ...
func (c Customer) GetStripeCustomerID() string {
	return c.StripeCustomerID
}
