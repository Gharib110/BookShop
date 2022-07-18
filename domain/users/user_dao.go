package users

import (
	"github.com/Gharib110/bookstore_users_api/database/mysql"
	"github.com/Gharib110/bookstore_users_api/utils/errors"
	"net/http"
)

func (user *User) Get(userId int64) *errors.RestErr {
	if err := mysql2.UserDB.Ping(); err != nil {
		return &errors.RestErr{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Error:   "StatusInternalServerError",
		}
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	return nil
}
