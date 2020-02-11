package usecase

import (
	"github.com/kindaidensan/UMR/domain"
	"math/rand"
	"time"
	"strconv"
	"errors"
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

func (interactor *AccountInteractor) FindTemporaryAccount(id string) (domain.Account, error) {
	account, err := interactor.accountRepository.FindByIdFromTemporary(id)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (interactor *AccountInteractor) AuthenticationTemporaryAccount(clientAuth domain.AuthenticationCode) error {
	serverAuth, err := interactor.authenticationCodeRepository.FindID(clientAuth.ID)	
	if err != nil {
		return err
	}
	if serverAuth.Code != clientAuth.Code {
		return errors.New("Faild: certification")
	}
	return nil
}

func (interactor *AccountInteractor ) Registration(account domain.Account)  error {
	err := interactor.accountRepository.Store(account)
	if err != nil {
		return err
	}

	return nil
}
