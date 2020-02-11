package database

import (
	"github.com/kindaidensan/UMR/domain" 
)

type AuthenticationCodeRepository struct {
	redisHandler RedisHandler
}

func NewAuthenticationCodeRepository(redisHandler RedisHandler) *AuthenticationCodeRepository {
	authenticationCodeRepository := AuthenticationCodeRepository {
		redisHandler: redisHandler,
	}
	return &authenticationCodeRepository
}

func (repo *AuthenticationCodeRepository) Store(authenticationCode domain.AuthenticationCode) error {
	err := repo.redisHandler.Set(authenticationCode.ID, authenticationCode.Code)
	if err != nil {
		return err
	}
	return nil
}

func (repo *AuthenticationCodeRepository) FindID(id string) (domain.AuthenticationCode, error) {
	code, err := repo.redisHandler.Get(id)
	if err != nil {
		return domain.AuthenticationCode{}, err
	}
	return domain.AuthenticationCode { ID: id, Code: code }, nil
}
