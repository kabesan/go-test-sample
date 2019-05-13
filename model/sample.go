package model

import (
	"errors"
)

var (
	ErrUserNotFound       = errors.New("user_not_found")
	ErrUserError          = errors.New("user_error")
	ErrCustomerNotFound   = errors.New("customer_not_found")
	ErrCustomerError      = errors.New("customer_error")
	ErrSubscriptionExists = errors.New("subscription_exists")
	ErrSubscriptionError  = errors.New("subscription_error")
)

type SampleModel struct {
	UserRepo         UserRepository
	CustomerRepo     CustomerRepository
	SubscriptionRepo SubscriptionRepository
}

func (sample *SampleModel) RegisterNewSubscription(userID, planID string) (*Subscription, error) {
	user, err := sample.UserRepo.ReadUser(&ReadUserParams{
		UserID: userID,
	})
	if err != nil {
		return nil, ErrUserError
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	if user.CustomerID == "" {
		return nil, ErrCustomerNotFound
	}

	cus, err := sample.CustomerRepo.ReadCustomer(&ReadCustomerParams{
		CustomerID: user.CustomerID,
	})
	if err != nil {
		return nil, ErrCustomerError
	}
	if cus == nil {
		return nil, ErrCustomerNotFound
	}
	if len(cus.Subscriptions) > 0 {
		return nil, ErrSubscriptionExists
	}

	newSub, err := sample.SubscriptionRepo.CreateSubscription(&CreateSubscriptionParams{
		CustomerID: cus.CustomerID,
		PlanID:     planID,
	})
	if err != nil {
		return nil, ErrSubscriptionError
	}

	return newSub, nil
}
