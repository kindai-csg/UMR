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

func (repo *AppRepository) FindByUserId(userId string) ([]domain.App, error) {
	query := "SELECT * FROM app " +
			 "WHERE " +
				 "owner = '" + userId + "';" 
	rows, err := repo.SqlHandler.Query(query)
	if err != nil {
		return nil, err
	}
	apps := []domain.App {}
	for rows.Next() {
		var id string
		var name string
		var description string
		var consumerkey string
		var consumersecret string
		var callback string
		var owner string
		if err := rows.Scan(&id, &name, &description, &consumerkey, &consumersecret, &callback, &owner); err != nil {
			return nil, err
		}
		app := domain.App{}
		app.ID = id
		app.Name = name
		app.Description = description
		app.ConsumerKey = consumerkey
		app.ConsumerSecret = consumersecret
		app.Callback = callback
		apps = append(apps, app)
	}
	return apps, nil
}

func (repo *AppRepository) FindByConsumerKey(consumerKey string) (domain.App, error) {
	query := "SELECT * FROM app " +
			 "WHERE " +
				 "consumerkey = '" + consumerKey + "';" 
	rows, err := repo.SqlHandler.Query(query)
	if err != nil {
		return domain.App{}, err
	}
	app := domain.App {}
	for rows.Next() {
		var id string
		var name string
		var description string
		var consumerkey string
		var consumersecret string
		var callback string
		var owner string
		if err := rows.Scan(&id, &name, &description, &consumerkey, &consumersecret, &callback, &owner); err != nil {
			return app, err
		}
		app.ID = id
		app.Name = name
		app.Description = description
		app.ConsumerKey = consumerkey
		app.ConsumerSecret = consumersecret
		app.Callback = callback
	}
	return app, nil
}

func (repo *AppRepository) RecreateKey(userId string, appId string, consumerKey string, consumerSecret string) (domain.App, error) {
	return domain.App{}, nil
}