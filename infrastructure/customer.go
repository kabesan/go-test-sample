package infrastructure

import (
	"github.com/kabesan/go-test-sample/model"
	"github.com/stripe/stripe-go/client"
)

type StripeCustomerRepository struct {
	StripeAPI *client.API
}

func (repo *StripeCustomerRepository) CreateCustomer(params *model.CreateCustomerParams) (*model.Customer, error) {
	return nil, nil
}

func (repo *StripeCustomerRepository) ReadCustomer(params *model.ReadCustomerParams) (*model.Customer, error) {
	cus, err := repo.StripeAPI.Customers.Get(params.CustomerID, nil)
	if err != nil {
		return nil, err
	}

	subs := make([]*model.Subscription, cus.Subscriptions.TotalCount)
	for i, sub := range cus.Subscriptions.Data {
		subs[i] = &model.Subscription{
			SubscriptionID: sub.ID,
		}
	}

	return &model.Customer{
		CustomerID:    cus.ID,
		Subscriptions: subs,
	}, nil
}

func (repo *StripeCustomerRepository) UpdateCustomer(params *model.UpdateCustomerParams) (*model.Customer, error) {
	return nil, nil
}

func (repo *StripeCustomerRepository) DeleteCustomer(params *model.DeleteCustomerParams) error {
	return nil
}
