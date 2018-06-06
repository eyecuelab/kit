package stripe

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"github.com/stripe/stripe-go/customer"
)

// Client stripe client data structure
type Client struct {
	Parent Parent
}

// PaymentSource payment source model interface
type PaymentSource interface {
	SetSource(string, interface{}) error
}

// Setup init stripe setup
func Setup(key string) error {
	if key == "" {
		return errors.New("Stripe secret key is not set")
	}
	stripe.Key = key
	return nil
}

// NewClient init new client
func NewClient(parent Parent) (*Client, error) {
	if err := Setup(viper.GetString("stripe_secret_key")); err != nil {
		return nil, err
	}

	return &Client{Parent: parent}, nil
}

// FetchCustomer find customer record
func (i *Client) FetchCustomer() (*stripe.Customer, error) {
	customerID := i.Parent.GetStripeCustomerID()
	if customerID == nil {
		return nil, nil
	}
	c, err := customer.Get(*customerID, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// EnsureCustomer find or create a customer record
func (i *Client) EnsureCustomer() error {
	c, err := i.FetchCustomer()
	if err != nil || c != nil {
		return err
	}

	params := &stripe.CustomerParams{
		Description: i.Parent.StripeCustomerDescription(),
	}

	c, err = customer.New(params)
	if err != nil {
		return err
	}

	return i.Parent.SetStripeCustomerID(c.ID)
}

// CreatePaymentSource create payment source
func (i *Client) CreatePaymentSource(paymentType string, token string, ps PaymentSource) error {
	if paymentType == "credit_card" {
		return i.CreateCreditCard(&token, ps)
	}

	return fmt.Errorf("Payment type '%s' is not supported", paymentType)
}

// CreateCreditCard create credit card payment source
func (i *Client) CreateCreditCard(token *string, ps PaymentSource) error {
	c, err := card.New(&stripe.CardParams{
		Customer: i.Parent.GetStripeCustomerID(),
		Token:    token,
	})
	if err != nil {
		return err
	}

	return ps.SetSource(c.ID, c)
}
