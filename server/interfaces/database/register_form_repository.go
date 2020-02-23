package database

import (
	"github.com/kindaidensan/UMR/domain" 
)

type RegisterFormRepository struct {
	RedisHandler RedisHandler
}

func NewRegisterFormRepository(redisHandler RedisHandler) *RegisterFormRepository {
	registerFormRepository := RegisterFormRepository {
		RedisHandler: redisHandler,
	}
	return &registerFormRepository
}

func (repo *RegisterFormRepository) Set(form domain.RegisterForm) error {
	err := repo.RedisHandler.ExpireSetKey("register_form", form.Token, form.Time)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RegisterFormRepository) Get() (domain.RegisterForm, error) {
	form := domain.RegisterForm {}
	token, err := repo.RedisHandler.Get("register_form")
	if err != nil {
		return form, err
	}
	ttl, err := repo.RedisHandler.GetTtl("register_form")
	if err != nil {
		return form, err
	}
	form.Token = token
	form.Time = ttl
	return form, nil
}
