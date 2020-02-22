package usecase

import (
	"github.com/kindaidensan/UMR/domain"
	"github.com/google/uuid"
)

type SettingInteractor struct {
	RegisterFormRepository RegisterFormRepository
}

func NewSettingInteractor(registerFormRepository RegisterFormRepository) *SettingInteractor {
	settingInteractor := SettingInteractor {
		RegisterFormRepository: registerFormRepository,
	}
	return &settingInteractor
}

func (interactor *SettingInteractor) IssueFormToken(time int) (domain.RegisterForm, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return domain.RegisterForm{}, err
	}
	registerForm := domain.RegisterForm {
		Token: uuid.String(),
		Time: time,
	}
	err = interactor.RegisterFormRepository.Set(registerForm)
	if err != nil {
		return registerForm, err
	}
	return registerForm, nil
} 

func (interactor *SettingInteractor) GetFormToken() (domain.RegisterForm, error) {
	form, err := interactor.RegisterFormRepository.Get()
	if err != nil {
		return form, err
	}
	return form, nil
}
