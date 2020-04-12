package usecase_test

import (
	mock "github.com/kindaidensan/UMR/test/mock_usecase"
	"github.com/kindaidensan/UMR/domain"
	"github.com/kindaidensan/UMR/usecase"
	"errors"
	"testing"
	"github.com/golang/mock/gomock"
)

func TestSettingFormToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockRegisterFormRepository(ctrl)
	interactor := usecase.NewSettingInteractor(m)

	m.EXPECT().Set(gomock.Any()).Return(errors.New("error"))
	_, err := interactor.IssueFormToken(60)
	if err == nil {
		t.Errorf("faild: expectation error")
	}

	m.EXPECT().Get().Return(domain.RegisterForm{}, errors.New("error"))
	_, err = interactor.GetFormToken()
	if err == nil {
		t.Errorf("faild: expectation error")
	}	
}