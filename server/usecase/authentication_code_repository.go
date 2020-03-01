package usecase

import "github.com/kindaidensan/UMR/domain" 

type AuthenticationCodeRepository interface {
	Store(domain.AuthenticationCode) (error)
	FindID(string) (domain.AuthenticationCode, error)
	IncFailureCount(string) (error)
}
