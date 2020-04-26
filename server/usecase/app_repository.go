package usecase

import "github.com/kindaidensan/UMR/domain" 

type AppRepository interface {
	Create(string, domain.App) (domain.App, error)
	Delete(string, string) (error)
	Update(string, domain.App) (domain.App, error)
	Get(string) ([]domain.App, error)
	RecreateKey(string, string, string, string) (domain.App, error)
}