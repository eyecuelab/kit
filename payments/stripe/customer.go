package stripe

// Customer ...
type Customer struct {
	StripeCustomerID *string
}

// Source ...
type Source struct {
	StripeSourceID *string
}

// Parent parent model interface
type Parent interface {
	StripeCustomerDescription() *string
	GetStripeCustomerID() *string
	SetStripeCustomerID(*string) error
}

// GetStripeCustomerID ...
func (c Customer) GetStripeCustomerID() *string {
	return c.StripeCustomerID
}

// SetSource ...
func (i *Source) SetSource(sourceID string, source interface{}) error {
	i.StripeSourceID = &sourceID

	return nil
}
