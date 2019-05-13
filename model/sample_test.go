package model

import (
	"errors"
	"testing"
)

type mockUserRepository struct {
	UserRepository
	readUserFunc func(params *ReadUserParams) (*User, error)
}

func (repo *mockUserRepository) ReadUser(params *ReadUserParams) (*User, error) {
	return repo.readUserFunc(params)
}

type mockCustomerRepository struct {
	CustomerRepository
	readCustomerFunc func(params *ReadCustomerParams) (*Customer, error)
}

func (repo mockCustomerRepository) ReadCustomer(params *ReadCustomerParams) (*Customer, error) {
	return repo.readCustomerFunc(params)
}

type mockSubscriptionRepository struct {
	SubscriptionRepository
	createSubscriptionFunc func(params *CreateSubscriptionParams) (*Subscription, error)
}

func (repo *mockSubscriptionRepository) CreateSubscription(params *CreateSubscriptionParams) (*Subscription, error) {
	return repo.createSubscriptionFunc(params)
}

func TestRegisterNewSubscription(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		model := SampleModel{
			UserRepo: &mockUserRepository{
				readUserFunc: func(params *ReadUserParams) (user *User, e error) {
					return &User{
						UserID:     params.UserID,
						CustomerID: "cus_dummy",
					}, nil
				},
			},
			CustomerRepo: &mockCustomerRepository{
				readCustomerFunc: func(params *ReadCustomerParams) (customer *Customer, e error) {
					return &Customer{
						CustomerID:    "cus_dummy",
						Subscriptions: []*Subscription{},
					}, nil
				},
			},
			SubscriptionRepo: &mockSubscriptionRepository{
				createSubscriptionFunc: func(params *CreateSubscriptionParams) (subscription *Subscription, e error) {
					return &Subscription{
						SubscriptionID: "sub_dummy",
					}, nil
				},
			},
		}

		result, err := model.RegisterNewSubscription("user_dummy", "plan_dummy")
		if err != nil {
			t.Error()
		}
		if result == nil {
			t.Error()
		}
	})
	t.Run("異常系_Userが存在しない", func(t *testing.T) {
		model := SampleModel{
			UserRepo: &mockUserRepository{
				readUserFunc: func(params *ReadUserParams) (user *User, e error) {
					return nil, nil
				},
			},
		}

		result, err := model.RegisterNewSubscription("user_dummy", "plan_dummy")
		if err != ErrUserNotFound {
			t.Error(err)
		}
		if result != nil {
			t.Error()
		}
	})
	t.Run("異常系_UserにCustomerが紐付いてない", func(t *testing.T) {
		model := SampleModel{
			UserRepo: &mockUserRepository{
				readUserFunc: func(params *ReadUserParams) (user *User, e error) {
					return &User{
						UserID:     params.UserID,
						CustomerID: "",
					}, nil
				},
			},
		}

		result, err := model.RegisterNewSubscription("user_dummy", "plan_dummy")
		if err != ErrCustomerNotFound {
			t.Error(err)
		}
		if result != nil {
			t.Error()
		}
	})
	t.Run("異常系_User取得エラー", func(t *testing.T) {
		model := SampleModel{
			UserRepo: &mockUserRepository{
				readUserFunc: func(params *ReadUserParams) (user *User, e error) {
					return nil, errors.New("")
				},
			},
		}

		result, err := model.RegisterNewSubscription("user_dummy", "plan_dummy")
		if err != ErrUserError {
			t.Error(err)
		}
		if result != nil {
			t.Error()
		}
	})
	t.Run("異常系_Userに紐付いたCustomerが存在しない", func(t *testing.T) {
		model := SampleModel{
			UserRepo: &mockUserRepository{
				readUserFunc: func(params *ReadUserParams) (user *User, e error) {
					return &User{
						UserID:     params.UserID,
						CustomerID: "cus_dummy",
					}, nil
				},
			},
			CustomerRepo: &mockCustomerRepository{
				readCustomerFunc: func(params *ReadCustomerParams) (customer *Customer, e error) {
					return nil, nil
				},
			},
		}

		result, err := model.RegisterNewSubscription("user_dummy", "plan_dummy")
		if err != ErrCustomerNotFound {
			t.Error(err)
		}
		if result != nil {
			t.Error()
		}
	})
	t.Run("異常系_Subscriptionがすでにある", func(t *testing.T) {
		model := SampleModel{
			UserRepo: &mockUserRepository{
				readUserFunc: func(params *ReadUserParams) (user *User, e error) {
					return &User{
						UserID:     params.UserID,
						CustomerID: "cus_dummy",
					}, nil
				},
			},
			CustomerRepo: &mockCustomerRepository{
				readCustomerFunc: func(params *ReadCustomerParams) (customer *Customer, e error) {
					return &Customer{
						CustomerID: "cus_dummy",
						Subscriptions: []*Subscription{
							{
								SubscriptionID: "sub_dummy",
							},
						},
					}, nil
				},
			},
		}

		result, err := model.RegisterNewSubscription("user_dummy", "plan_dummy")
		if err != ErrSubscriptionExists {
			t.Error(err)
		}
		if result != nil {
			t.Error()
		}
	})
	t.Run("異常系_Customer取得エラー", func(t *testing.T) {
		model := SampleModel{
			UserRepo: &mockUserRepository{
				readUserFunc: func(params *ReadUserParams) (user *User, e error) {
					return &User{
						UserID:     params.UserID,
						CustomerID: "cus_dummy",
					}, nil
				},
			},
			CustomerRepo: &mockCustomerRepository{
				readCustomerFunc: func(params *ReadCustomerParams) (customer *Customer, e error) {
					return nil, errors.New("")
				},
			},
		}

		result, err := model.RegisterNewSubscription("user_dummy", "plan_dummy")
		if err != ErrCustomerError {
			t.Error(err)
		}
		if result != nil {
			t.Error()
		}
	})
	t.Run("異常系_Subscription登録エラー", func(t *testing.T) {
		model := SampleModel{
			UserRepo: &mockUserRepository{
				readUserFunc: func(params *ReadUserParams) (user *User, e error) {
					return &User{
						UserID:     params.UserID,
						CustomerID: "cus_dummy",
					}, nil
				},
			},
			CustomerRepo: &mockCustomerRepository{
				readCustomerFunc: func(params *ReadCustomerParams) (customer *Customer, e error) {
					return &Customer{
						CustomerID:    "cus_dummy",
						Subscriptions: []*Subscription{},
					}, nil
				},
			},
			SubscriptionRepo: &mockSubscriptionRepository{
				createSubscriptionFunc: func(params *CreateSubscriptionParams) (subscription *Subscription, e error) {
					return nil, errors.New("")
				},
			},
		}

		result, err := model.RegisterNewSubscription("user_dummy", "plan_dummy")
		if err != ErrSubscriptionError {
			t.Error(err)
		}
		if result != nil {
			t.Error()
		}
	})
}
