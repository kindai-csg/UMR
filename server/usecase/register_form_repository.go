package usecase

import "github.com/kindaidensan/UMR/domain"

type RegisterFormRepository interface {
	Set(domain.RegisterForm) (error)
	Get() (domain.RegisterForm, error)
}
