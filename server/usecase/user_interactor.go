package usecase

import (
	"github.com/kindaidensan/UMR/domain"
	"math/rand"
	"time"
)

type UserInteractor struct {
	accountRepository AccountRepository
	authenticationCodeRepository AuthenticationCodeRepository
}

func NewUserInteractor(accountRepository AccountRepository, authenticationCodeRepository AuthenticationCodeRepository) *UserInteractor {
	userInteractor := UserInteractor {
		accountRepository: accountRepository,
		authenticationCodeRepository: authenticationCodeRepository,
	}
	return &userInteractor
}

func (interactor *UserInteractor ) TemporaryRegistration(account domain.Account) error {
	err := interactor.accountRepository.TemporaryStore(account)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9000) + 1000 
	authentication := domain.AuthenticationCode {
		ID: account.ID,
		Code: code,
	}
	err = interactor.authenticationCodeRepository.Store(authentication)
	if err != nil {
		return err
	}
	return err
}

func (interactor *UserInteractor ) Registration(id string)  error {
	account, err := interactor.accountRepository.FindByIdFromTemporary(id)
	if err != nil {
		return err
	} 

	err = interactor.accountRepository.Store(account)
	if err != nil {
		return err
	}

	return nil
}
