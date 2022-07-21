package services

import (
	"github.com/Gharib110/BookShop/domain/users"
	mycrypto "github.com/Gharib110/BookShop/utils/crypto"
	date_utils "github.com/Gharib110/BookShop/utils/date"
	"github.com/Gharib110/BookShop/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(*users.User) (*users.User, *errors.RestErr)
	UpdateUser(*users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}
	if err := result.Get(userId); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) CreateUser(user *users.User) (*users.User, *errors.RestErr) {
	// TODO: The Logic should be implemented
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.CreatedAt = date_utils.GetNowString()
	user.Password = mycrypto.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) UpdateUser(user *users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.ID)
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

func (s *usersService) DeleteUser(id int64) *errors.RestErr {
	user := &users.User{ID: id}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: mycrypto.GetMD5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
