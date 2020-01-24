package usecase

import "github.com/kindaidensan/UMR/domain" 

type RegularAccountRepository interface {
	TemporaryStore(domain.RegularAccount) (error)
	FindByIdFromTemporary(string) (domain.RegularAccount, error)
	Store(domain.RegularAccount) (error)	
}
