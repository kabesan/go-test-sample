package infrastructure

import (
	"github.com/kabesan/go-test-sample/model"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

type StripeSubscriptionRepository struct {
	StripeAPI *client.API
}

func (repo *StripeSubscriptionRepository) CreateSubscription(params *model.CreateSubscriptionParams) (*model.Subscription, error) {
	sub, err := repo.StripeAPI.Subscriptions.New(&stripe.SubscriptionParams{
		Customer: stripe.String(params.CustomerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String(params.PlanID),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &model.Subscription{
		SubscriptionID: sub.ID,
	}, nil
}

func (repo *StripeSubscriptionRepository) ReadSubscription(params *model.ReadSubscriptionParams) (*model.Subscription, error) {
	return nil, nil
}

func (repo *StripeSubscriptionRepository) UpdateSubscription(params *model.UpdateSubscriptionParams) (*model.Subscription, error) {
	return nil, nil
}

func (repo *StripeSubscriptionRepository) DeleteSubscription(params *model.DeleteSubscriptionParams) error {
	return nil
}
