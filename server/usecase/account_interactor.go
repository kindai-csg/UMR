package usecase

import (
	"github.com/kindaidensan/UMR/domain"
	"math/rand"
	"time"
	"strconv"
	"errors"
)

type AccountInteractor struct {
	AccountRepository AccountRepository
	AuthenticationCodeRepository AuthenticationCodeRepository
}

func NewAccountInteractor(accountRepository AccountRepository, authenticationCodeRepository AuthenticationCodeRepository) *AccountInteractor {
	accountInteractor :=AccountInteractor {
		AccountRepository: accountRepository,
		AuthenticationCodeRepository: authenticationCodeRepository,
	}
	return &accountInteractor
}

func (interactor *AccountInteractor ) TemporaryRegistration(account domain.Account) error {
	err := interactor.AccountRepository.TemporaryStore(account)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9000) + 1000 
	authentication := domain.AuthenticationCode {
		ID: account.ID,
		Code: strconv.Itoa(code),
	}
	err = interactor.AuthenticationCodeRepository.Store(authentication)
	if err != nil {
		return err
	}
	return err
}

func (interactor *AccountInteractor) FindTemporaryAccount(id string) (domain.Account, error) {
	account, err := interactor.AccountRepository.FindByIdFromTemporary(id)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (interactor *AccountInteractor) AuthenticationTemporaryAccount(clientAuth domain.AuthenticationCode) error {
	serverAuth, err := interactor.AuthenticationCodeRepository.FindID(clientAuth.ID)	
	if err != nil {
		return err
	}
	if serverAuth.Code != clientAuth.Code {
		return errors.New("Faild: certification")
	}
	return nil
}

func (interactor *AccountInteractor ) Registration(account domain.Account)  error {
	account.UserIdNumber = account.StudentNumber[7:]+account.StudentNumber[1:2]+account.StudentNumber[5:6]
	account.GroupIdNumber = "1002"
	account.HomeDirectory = "/home/"+account.ID
	err := interactor.AccountRepository.Store(account)
	if err != nil {
		return err
	}

	return nil
}

func (interactor *AccountInteractor) DuplicateCheck(id string) error {
	ids, err := interactor.AccountRepository.GetAllUserID()
	if err != nil {
		return err
	}
	for _, alreay_id := range ids {
		if id == alreay_id {
			return errors.New("既にこのIDは使用されています.")
		}
	}
	return nil
}