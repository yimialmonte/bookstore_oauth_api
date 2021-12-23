package access_token

import (
	"strings"

	"github.com/yimialmonte/src/utils/errors"
)

// Repository ...
type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

// Service ...
type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetById(accessToken string) (*AccessToken, *errors.RestErr) {
	accessToken = strings.TrimSpace(accessToken)
	if len(accessToken) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	return s.repository.GetById(accessToken)
}
