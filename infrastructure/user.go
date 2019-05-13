package infrastructure

import (
	"database/sql"

	"github.com/kabesan/go-test-sample/model"
)

type RDBUserRepository struct {
	DB *sql.DB
}

func (repo *RDBUserRepository) CreateUser(params *model.CreateUserParams) (*model.User, error) {
	return nil, nil
}

func (repo *RDBUserRepository) ReadUser(params *model.ReadUserParams) (*model.User, error) {
	rows, err := repo.DB.Query(
		"SELECT u.user_id, uc.customer_id "+
			"FROM user u LEFT OUTER JOIN user_customer uc USING(user_id) "+
			"WHERE u.user_id = ?",
		params.UserID,
	)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}

	var user model.User
	err = rows.Scan(
		&user.UserID,
		&user.CustomerID,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *RDBUserRepository) UpdateUser(params *model.UpdateUserParams) (*model.User, error) {
	return nil, nil
}

func (repo *RDBUserRepository) DeleteUser(params *model.DeleteUserParams) error {
	return nil
}
