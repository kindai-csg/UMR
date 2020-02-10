package usecase

import (
	"github.com/kindaidensan/UMR/domain"
	"math/rand"
	"time"
	"strconv"
)

type AccountInteractor struct {
	accountRepository AccountRepository
	authenticationCodeRepository AuthenticationCodeRepository
}

func NewAccountInteractor(accountRepository AccountRepository, authenticationCodeRepository AuthenticationCodeRepository) *AccountInteractor {
	accountInteractor :=AccountInteractor {
		accountRepository: accountRepository,
		authenticationCodeRepository: authenticationCodeRepository,
	}
	return &accountInteractor
}

func (interactor *AccountInteractor ) TemporaryRegistration(account domain.Account) error {
	err := interactor.accountRepository.TemporaryStore(account)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9000) + 1000 
	authentication := domain.AuthenticationCode {
		ID: account.ID,
		Code: strconv.Itoa(code),
	}
	err = interactor.authenticationCodeRepository.Store(authentication)
	if err != nil {
		return err
	}
	return err
}

func (interactor *AccountInteractor ) Registration(account domain.Account)  error {
	err := interactor.accountRepository.Store(account)
	if err != nil {
		return err
	}

	return nil
}
