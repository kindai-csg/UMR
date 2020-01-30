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
	userInteractor  := usecase.NewUserInteractor(aMock, authMock) 

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


	aMock.EXPECT().FindByIdFromTemporary("id").Return(account, nil)
	aMock.EXPECT().Store(account).Return(nil)
	if (nil != userInteractor .Registration("id")) {
		t.Errorf("faild: Registration / Expectation: return nil")
	}

	aMock.EXPECT().FindByIdFromTemporary("id").Return(account, err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userInteractor .Registration("id"))) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return type err")
	}

	aMock.EXPECT().FindByIdFromTemporary("id").Return(account, nil)
	aMock.EXPECT().Store(account).Return(err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userInteractor .Registration("id"))) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return type err")
	}
}