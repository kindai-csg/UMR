package usecase

import (
	"errors"
)

type AuthenticationInteractor struct {
	RegisterFormRepository RegisterFormRepository
}

func NewAuthenticationInteractor(registerFormRepository RegisterFormRepository) *AuthenticationInteractor {
	authenticationInteractor := AuthenticationInteractor {
		RegisterFormRepository: registerFormRepository,
	}
	return &authenticationInteractor
}

func (interactor *AuthenticationInteractor) AuthenticationFormToken(token string) error {
	form, err := interactor.RegisterFormRepository.Get()
	if err != nil {
		return err
	}
	if form.Token != token {
		return errors.New("faild: token authentication")
	} 
	return nil
}
