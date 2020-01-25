package usecase

import (
	"github.com/kindaidensan/UMR/domain"
	"math/rand"
	"time"
)

type UserRegistrationInteractor struct {
	accountRepository AccountRepository
	regularAccountRepository RegularAccountRepository
}

func NewUserRegistrationInteractor(accountRepository AccountRepository, regularAccountRepository RegularAccountRepository) *UserRegistrationInteractor {
	userRegistrationInteractor := UserRegistrationInteractor{accountRepository: accountRepository, regularAccountRepository: regularAccountRepository}
	return &userRegistrationInteractor
}

func (interactor *UserRegistrationInteractor ) TemporaryRegistration(account domain.Account, regular domain.RegularAccount) (int, error) {
	err := interactor.accountRepository.TemporaryStore(account)
	if err != nil {
		return -1, err
	}
	err = interactor.regularAccountRepository.TemporaryStore(regular)
	if  err != nil {
		return -1, err
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000, err
}

func (interactor *UserRegistrationInteractor ) Registration(id string)  error {
	account, err := interactor.accountRepository.FindByIdFromTemporary(id)
	if err != nil {
		return err
	} 

	regular, err := interactor.regularAccountRepository.FindByIdFromTemporary(id)
	if err != nil {
		return err
	}

	err = interactor.accountRepository.Store(account)
	if err != nil {
		return err
	}

	err = interactor.regularAccountRepository.Store(regular)
	if err != nil {
		return err
	}
	return nil
}
