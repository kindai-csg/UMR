package usecase

import (
	"github.com/kindaidensan/UMR/domain"
)

type AppInteractor struct {
	AppRepository AppRepository
}

func NewAppInteractor(appRepository AppRepository) *AppInteractor {
	appInteractor := AppInteractor {
		AppRepository: appRepository,
	}
	return &appInteractor
}

func (interactor *AppInteractor) CreateApplication(userId string, app domain.App) (domain.App, error) {
	app.ConsumerKey = RandString(32)
	app.ConsumerSecret = RandString(32)
	a, err := interactor.AppRepository.Create(userId, app)
	if err != nil {
		return domain.App{}, err
	}
	return a, nil
}

func (interactor *AppInteractor) DeleteApplication(userId string, appId string) error {
	err := interactor.AppRepository.Delete(userId, appId)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *AppInteractor) UpdateApplication(userId string, app domain.App) (domain.App, error) {
	a, err := interactor.AppRepository.Update(userId, app)
	if err != nil {
		return domain.App{}, err
	}
	return a, nil
}

func (interactor *AppInteractor) GetApplication(userId string) ([]domain.App, error) {
	apps, err := interactor.AppRepository.Get(userId)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (interactor *AppInteractor) RecreatehKey(userId string, appId string) (domain.App, error) {
	consumerKey := RandString(32)
	consumerSecret := RandString(32)
	app, err := interactor.AppRepository.RecreateKey(userId, appId, consumerKey, consumerSecret)
	if err != nil {
		return domain.App{}, err
	}
	return app, nil
}
