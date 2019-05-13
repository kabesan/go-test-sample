package model

type Customer struct {
	CustomerID    string
	Subscriptions []*Subscription
}

type CreateCustomerParams struct {
}

type ReadCustomerParams struct {
	CustomerID string
}

type UpdateCustomerParams struct {
}

type DeleteCustomerParams struct {
}

type CustomerRepository interface {
	CreateCustomer(params *CreateCustomerParams) (*Customer, error)
	ReadCustomer(params *ReadCustomerParams) (*Customer, error)
	UpdateCustomer(params *UpdateCustomerParams) (*Customer, error)
	DeleteCustomer(params *DeleteCustomerParams) error
}
