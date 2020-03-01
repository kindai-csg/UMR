package database

import (
	"github.com/kindaidensan/UMR/domain" 
)

type AuthenticationCodeRepository struct {
	RedisHandler RedisHandler
}

func NewAuthenticationCodeRepository(redisHandler RedisHandler) *AuthenticationCodeRepository {
	authenticationCodeRepository := AuthenticationCodeRepository {
		RedisHandler: redisHandler,
	}
	return &authenticationCodeRepository
}

func (repo *AuthenticationCodeRepository) Store(authenticationCode domain.AuthenticationCode) error {
	err := repo.RedisHandler.Set("auth_"+authenticationCode.ID, authenticationCode.Code)
	if err != nil {
		return err
	}
	return nil
}

func (repo *AuthenticationCodeRepository) FindID(id string) (domain.AuthenticationCode, error) {
	code, err := repo.RedisHandler.Get("auth_"+id)
	if err != nil {
		return domain.AuthenticationCode{}, err
	}
	return domain.AuthenticationCode { ID: id, Code: code }, nil
}
