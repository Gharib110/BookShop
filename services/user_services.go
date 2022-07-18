package services

import (
	"github.com/Gharib110/bookstore_users_api/domain/users"
	"github.com/Gharib110/bookstore_users_api/utils/errors"
)

func CreateUser(user *users.User) (*users.User, *errors.RestErr) {
	// TODO: The Logic should be implemented
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}
