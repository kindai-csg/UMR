package database

import (
	"github.com/kindaidensan/UMR/domain" 
	"errors"
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

func (repo *AuthenticationCodeRepository) IncFailureCount(id string) error {
	count, err := repo.RedisHandler.Incr("count_"+id)
	if err != nil {
		return err
	}
	if count < 3 {
		return nil
	}
	err = repo.RedisHandler.MultiDel([]string { "tmp_"+id, "auth_"+id, "count_"+id })
	if err != nil {
		return err
	}
	return errors.New("試行回数を超えました. もう一度登録からお願いします.")
}
