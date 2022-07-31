package acc_token

import "github.com/Gharib110/BookShop/BookShop/utils/errors"

type AccessTokenService interface {
	GetByID() (*AccessToken, *errors.RestErr)
}

type AccTokenService struct {
}

func NewAccTokenService() AccessTokenService {
	return &AccTokenService{}
}

func (a *AccTokenService) GetByID() (*AccessToken, *errors.RestErr) {
	return nil, nil
}
