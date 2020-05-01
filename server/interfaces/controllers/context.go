package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	PostForm(string) string
	Status(int)
	JSON(int, interface{})
	Get(string) (interface{}, bool)
}
