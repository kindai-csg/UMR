package usecase

import (
	"github.com/kindaidensan/UMR/domain"
	"math/rand"
	"time"
	"strconv"
	"errors"
	"crypto/md5"
	"encoding/hex"
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

func (interactor *AccountInteractor ) TemporaryRegistration(account domain.Account) (string, error) {
	err := interactor.AccountRepository.TemporaryStore(account)
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(90000000) + 10000000 
	authentication := domain.AuthenticationCode {
		ID: account.ID,
		Code: strconv.Itoa(code),
	}
	err = interactor.AuthenticationCodeRepository.Store(authentication)
	if err != nil {
		return "", err
	}
	return authentication.Code, err
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
	if serverAuth.Code == clientAuth.Code {
		err = interactor.AuthenticationCodeRepository.DeleteAuthData(clientAuth.ID)
		if err != nil {
			return err
		}
		return nil
	}
	err = interactor.AuthenticationCodeRepository.IncFailureCount(clientAuth.ID)
	if err != nil {
		return err
	}
	return errors.New("認証に失敗しました.")
}

func (interactor *AccountInteractor ) Registration(account domain.Account)  error {
	uidHead := "5"
	if (account.StudentNumber[7:8] != "0") {
		uidHead = account.StudentNumber[7:8]
	}
	account.UserIdNumber = uidHead + account.StudentNumber[8:]+account.StudentNumber[1:2]+account.StudentNumber[5:6]
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

func (interactor *AccountInteractor) GetAllAccounts() ([]domain.Account, error) {
	accounts, err := interactor.AccountRepository.GetAllAccounts()
	if err != nil {
		return []domain.Account {}, err
	}
	return accounts, nil
}

func (interactor *AccountInteractor) GetAllNonActiveAccountID() ([]domain.Account, error) {
	ids, err := interactor.AccountRepository.GetAllNonActiveAccountID()
	if err != nil {
		return nil, err
	}
	accounts := []domain.Account {}
	for _, id := range ids {
		accounts = append(accounts, domain.Account { ID: id })
	}
	return accounts, nil
}

func (interactor *AccountInteractor) AuthenticationCheck(id string) (error) {
	ids, err := interactor.AccountRepository.GetAllNonActiveAccountID()
	if err != nil {
		return err
	}
	for _, _id := range ids {
		if id == _id {
			return nil
		} 
	}
	return errors.New("認証が完了していないアカウントです")
}

func (interactor *AccountInteractor) DeleteAccount(id string) (error) {
	err := interactor.AccountRepository.DeleteAccount(id)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *AccountInteractor) AuthenticationAdminAccount(account domain.AdminAccount) (error) {
	accounts, err := interactor.AccountRepository.GetAdminAccounts()
	if err != nil {
		return err
	}
	for _, a := range accounts {
		md5 := md5.Sum([]byte(account.Password))
		if a.ID == account.ID && a.Password == hex.EncodeToString(md5[:]) {
			return nil
		}
	}
	return errors.New("IDかPasswordが間違っています")
}
