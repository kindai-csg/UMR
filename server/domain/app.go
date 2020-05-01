package domain

type App struct {
	ID string `form:"id"`
	Name string `form:"name"`
	Description string `form:"description"`
	ConsumerKey string `form:"consumerkey"`
	ConsumerSecret string `form:"consumersecret"`
	Callback string `form:"callback"`
}