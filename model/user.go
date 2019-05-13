package model

type User struct {
	UserID     string
	CustomerID string
}

type CreateUserParams struct {
}

type ReadUserParams struct {
	UserID string
}

type UpdateUserParams struct {
}

type DeleteUserParams struct {
}

type UserRepository interface {
	CreateUser(params *CreateUserParams) (*User, error)
	ReadUser(params *ReadUserParams) (*User, error)
	UpdateUser(params *UpdateUserParams) (*User, error)
	DeleteUser(params *DeleteUserParams) error
}
