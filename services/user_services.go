package services

import (
	"github.com/Gharib110/bookstore_users_api/domain/users"
	"github.com/Gharib110/bookstore_users_api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}
	if err := result.Get(userId); err != nil {
		return nil, err
	}

	return result, nil
}

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

func UpdateUser(user *users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	err = current.Update()
	if err != nil {
		return nil, err
	}
	return current, nil
}
