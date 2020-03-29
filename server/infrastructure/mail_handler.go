package infrastructure

import (
	"net/smtp"
)

type MailHandler struct {
	from string
	host string
	auth smtp.Auth
}

func NewMailHandler() *MailHandler {
	mailHandler := MailHandler {
		from: "admin@kindai-csg.dev",
		host: "smtp",
		auth: smtp.PlainAuth("", "admin", "password", "smtp"),
	}
	return &mailHandler
}

func (handler *MailHandler) SendMail(to string, subject string, body string) error {
	msg := []byte(
		"From: " + handler.from + "\r\n" + 
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	err := smtp.SendMail(handler.host + ":1025", nil, handler.from, []string{ to }, msg)
	if  handler.auth == nil {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
