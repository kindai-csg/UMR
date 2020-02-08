package usecase_test

import (
	mock "github.com/kindaidensan/UMR/test/mock_interfaces/mock_database"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/interfaces/database"
	"errors"
	"testing"
	"github.com/golang/mock/gomock"
)

func NewMockInstance(t *testing.T) (*database.AccountRepository, *mock.MockLdapHandler, *mock.MockRedisHandler) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()	
	ldap := mock.NewMockLdapHandler(ctrl)
	redis := mock.NewMockRedisHandler(ctrl) 
	repository := database.NewAccountRepository(ldap, redis)
	return repository, ldap, redis
}

func TestTemporaryStore(t *testing.T) {
	repository, _, redis := NewMockInstance(t)
	account := domain.Account {
		ID: "test",
		Password: "test",
		Name: "test",
		EmailAddress: "test",
		StudentNumber: "test",
		AccountType: "test",
	}
	redis.EXPECT().RPush(gomock.Any(), gomock.Any()).Return(nil)
	err := repository.TemporaryStore(account)
	if err != nil {
		t.Errorf("faild: Expectation: return nil")
	}

	redis.EXPECT().RPush(gomock.Any(), gomock.Any()).Return(errors.New("error"))
	err = repository.TemporaryStore(account)
	if err == nil {
		t.Errorf("faild: Expectation: return err")
	}
}

func TestFindByIdFromTemporary(t *testing.T) {
	repository, _, redis := NewMockInstance(t)
	id := "test"
	redis.EXPECT().LPop(id, 5).Return([]string{}, errors.New("error"))	
	_, err := repository.FindByIdFromTemporary(id)
	if err == nil {
		t.Errorf("faild: Expectation: return err")	
	}
	
	account_array := []string{"test", "pass", "name", "mail", "snumber", "type"}
	redis.EXPECT().LPop(id, 5).Return(account_array, nil)
	account, err := repository.FindByIdFromTemporary(id)
	if err != nil {
		t.Errorf("faild: Expectation: return nil")	
	}	
	if id != account.ID {
		t.Errorf("faild: does not much ID")	
	}
	if account_array[0] != account.Password {
		t.Errorf("faild: does not much Password")	
	}
	if account_array[1] != account.Name {
		t.Errorf("faild: does not much Name")	
	}
	if account_array[2] != account.EmailAddress {
		t.Errorf("faild: does not much EmailAddress")	
	}
	if account_array[3] != account.StudentNumber {
		t.Errorf("faild: does not much StudentNumber")	
	}
	if account_array[4] != account.AccountType {
		t.Errorf("faild: does not much AccountType")	
	}

}

func TestStore(t *testing.T) {
	repository, ldap, _ := NewMockInstance(t)
	account := domain.Account {
		ID: "test",
		Password: "test",
		Name: "test",
		EmailAddress: "test",
		StudentNumber: "2010400200",
		AccountType: "Regular",
	}
	ldap.EXPECT().AddRequest(gomock.Any()).Return(nil)
	err := repository.Store(account)
	if err != nil {
		t.Errorf("faild: Expectation: return nil")
	}
	ldap.EXPECT().AddRequest(gomock.Any()).Return(errors.New("error"))
	err = repository.Store(account)
	if err == nil {
		t.Errorf("faild: Expectation: return error")
	}
}
