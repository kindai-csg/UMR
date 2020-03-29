package controllers

type MailHandler interface {
	SendMail(string, string, string) (error)
}
