package model

type Subscription struct {
	SubscriptionID string
}

type CreateSubscriptionParams struct {
	CustomerID string
	PlanID     string
}

type ReadSubscriptionParams struct {
}

type UpdateSubscriptionParams struct {
}

type DeleteSubscriptionParams struct {
}

type SubscriptionRepository interface {
	CreateSubscription(params *CreateSubscriptionParams) (*Subscription, error)
	ReadSubscription(params *ReadSubscriptionParams) (*Subscription, error)
	UpdateSubscription(params *UpdateSubscriptionParams) (*Subscription, error)
	DeleteSubscription(params *DeleteSubscriptionParams) error
}
