package test_usecase

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
	rMock := mock.NewMockRegularAccountRepository(ctrl) 

	account := domain.Account{} 
	regular := domain.RegularAccount{}
	userRegistrationInteractor := usecase.NewUserRegistrationInteractor(aMock, rMock) 

	aMock.EXPECT().TemporaryStore(account).Return(nil)
	rMock.EXPECT().TemporaryStore(regular).Return(nil)
	if (nil != userRegistrationInteractor.TemporaryRegistration(account, regular)) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return nil")
	}

	err := errors.New("error")

	aMock.EXPECT().TemporaryStore(account).Return(err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userRegistrationInteractor.TemporaryRegistration(account, regular))) {
		t.Errorf("faild: TemporaryRegistraton / Expectation: return type err")
	}

	aMock.EXPECT().TemporaryStore(account).Return(nil)
	rMock.EXPECT().TemporaryStore(regular).Return(err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userRegistrationInteractor.TemporaryRegistration(account, regular))) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return type err")
	}

	aMock.EXPECT().FindByIdFromTemporary("id").Return(account, nil)
	rMock.EXPECT().FindByIdFromTemporary("id").Return(regular, nil)
	aMock.EXPECT().Store(account).Return(nil)
	rMock.EXPECT().Store(regular).Return(nil)
	if (nil != userRegistrationInteractor.Registration("id")) {
		t.Errorf("faild: Registration / Expectation: return nil")
	}

	aMock.EXPECT().FindByIdFromTemporary("id").Return(account, err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userRegistrationInteractor.Registration("id"))) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return type err")
	}

	aMock.EXPECT().FindByIdFromTemporary("id").Return(account, nil)
	rMock.EXPECT().FindByIdFromTemporary("id").Return(regular, err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userRegistrationInteractor.Registration("id"))) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return type err")
	}

	aMock.EXPECT().FindByIdFromTemporary("id").Return(account, nil)
	rMock.EXPECT().FindByIdFromTemporary("id").Return(regular, nil)
	aMock.EXPECT().Store(account).Return(err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userRegistrationInteractor.Registration("id"))) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return type err")
	}

	aMock.EXPECT().FindByIdFromTemporary("id").Return(account, nil)
	rMock.EXPECT().FindByIdFromTemporary("id").Return(regular, nil)
	aMock.EXPECT().Store(account).Return(nil)
	rMock.EXPECT().Store(regular).Return(err)
	if (reflect.TypeOf(errors.New("")) != reflect.TypeOf(userRegistrationInteractor.Registration("id"))) {
		t.Errorf("faild: TemporaryRegistration / Expectation: return type err")
	}
}