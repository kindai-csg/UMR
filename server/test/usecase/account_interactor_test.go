package usecase_test

import (
	mock "github.com/kindaidensan/UMR/test/mock_usecase"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/usecase"
	"errors"
	"reflect"
	"testing"
	"github.com/golang/mock/gomock"
)

func TestRegistration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	aMock := mock.NewMockAccountRepository(ctrl)	
	authMock := mock.NewMockAuthenticationCodeRepository(ctrl)

	account := domain.Account{} 
	userInteractor  := usecase.NewAccountInteractor(aMock, authMock) 

	aMock.EXPECT().TemporaryStore(account).Return(nil)
	authMock.EXPECT().Store(gomock.Any()).Return(nil)
	e := userInteractor .TemporaryRegistration(account) 
	if (nil != e) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return nil")
	}

	err := errors.New("error")

	aMock.EXPECT().TemporaryStore(account).Return(nil)
	authMock.EXPECT().Store(gomock.Any()).Return(err)
	e = userInteractor .TemporaryRegistration(account) 
	if (nil == e) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return nil")
	}

	aMock.EXPECT().TemporaryStore(account).Return(err)
	e = userInteractor .TemporaryRegistration(account) 
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(e)) {
		t.Errorf("faild: TemporaryRegistraton / Expectation: return type err")
	}

	aMock.EXPECT().Store(account).Return(nil)
	if (nil != userInteractor .Registration(account)) {
		t.Errorf("faild: Registration / Expectation: return nil")
	}

	aMock.EXPECT().Store(account).Return(err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userInteractor .Registration(account))) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return type err")
	}
}

func TestFindTemporaryAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	aMock := mock.NewMockAccountRepository(ctrl)	
	authMock := mock.NewMockAuthenticationCodeRepository(ctrl)
	userInteractor  := usecase.NewAccountInteractor(aMock, authMock) 

	aMock.EXPECT().FindByIdFromTemporary(gomock.Any()).Return(domain.Account{}, nil)
	_, err := userInteractor.FindTemporaryAccount("test")
	if err != nil {
		t.Errorf("faild: Expectation: return nil")	
	}

	aMock.EXPECT().FindByIdFromTemporary(gomock.Any()).Return(domain.Account{}, errors.New("error"))
	_, err = userInteractor.FindTemporaryAccount("test")
	if err == nil {
		t.Errorf("faild: Expectation: return err")	
	}
}

func TestAuthenticationTemporaryAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	aMock := mock.NewMockAccountRepository(ctrl)	
	authMock := mock.NewMockAuthenticationCodeRepository(ctrl)
	userInteractor  := usecase.NewAccountInteractor(aMock, authMock) 	

	client_auth := domain.AuthenticationCode {
		ID: "test",
		Code: "3340",
	}
	authMock.EXPECT().FindID("test").Return(client_auth, nil)
	err := userInteractor.AuthenticationTemporaryAccount(client_auth)
	if err != nil {
		t.Errorf("faild: Expectation: return nil")	
	}
	authMock.EXPECT().FindID("test").Return(client_auth, errors.New("error"))
	err = userInteractor.AuthenticationTemporaryAccount(client_auth)
	if err == nil {
		t.Errorf("faild: Expectation: return err")	
	}
	authMock.EXPECT().FindID("test").Return(domain.AuthenticationCode{ID: "test", Code: "0334"}, nil)
	err = userInteractor.AuthenticationTemporaryAccount(client_auth)
	if err == nil {
		t.Errorf("faild: Expectation: return err")	
	}

}