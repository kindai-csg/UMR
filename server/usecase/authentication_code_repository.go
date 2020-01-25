package usecase

import "github.com/kindaidensan/UMR/domain" 

type AuthenticationCodeRepository interface {
	Store(domain.AuthenticationCode) (error)
	FindAll() (domain.AuthenticationCode, error)
}
