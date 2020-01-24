package usecase

import "github.com/kindaidensan/UMR/domain" 

type AccountRepository interface {
	TemporaryStore(domain.Account) (error)
	FindByIdFromTemporary(string) (domain.Account, error)
	Store(domain.Account) (error)	
}
