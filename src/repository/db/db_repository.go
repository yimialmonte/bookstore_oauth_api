package db

import (
	"github.com/yimialmonte/src/domain/access_token"
	"github.com/yimialmonte/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

func New() DbRepository {
	return &dbRepository{}
}

type dbRepository struct{}

func (d *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
