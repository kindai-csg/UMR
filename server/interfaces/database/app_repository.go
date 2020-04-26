package database

import (
	"github.com/kindaidensan/UMR/domain" 
)

type AppRepository struct {
	SqlHandler SqlHandler
}

func NewAppRepository(sqlHandler SqlHandler) *AppRepository {
	appRepository := AppRepository {
		SqlHandler: sqlHandler,
	}
	return &appRepository
}

func (repo *AppRepository) Create(userId string, app domain.App) (domain.App, error) {
	query := "INSERT INTO app " +
			 "VALUES(" +
				 "'"+app.ID + "'," +
				 "'"+app.Name + "'," +
				 "'"+app.Description + "'," +
				 "'"+app.ConsumerKey + "'," +
				 "'"+app.ConsumerSecret + "'," +
				 "'"+app.Callback + "'," +
				 "'"+userId  + "'" +
			 ");"
	_, err := repo.SqlHandler.Query(query)
	if err != nil {
		return domain.App{}, err
	}
	return app, nil
}

func (repo *AppRepository) Delete(userId string, appId string) error {
	return nil
}

func (repo *AppRepository) Update(userId string, app domain.App) (domain.App, error) {
	return domain.App{}, nil
}

func (repo *AppRepository) Get(userId string) ([]domain.App, error) {
	return nil, nil 
}

func (repo *AppRepository) RecreateKey(userId string, appId string, consumerKey string, consumerSecret string) (domain.App, error) {
	return domain.App{}, nil
}