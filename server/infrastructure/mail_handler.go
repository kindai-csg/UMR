package infrastructure

import (
	"net/smtp"
	"log"
)

type MailHandler struct {
	from string
	host string
	auth smtp.Auth
}

type MailConfig struct {
	From string  `toml:"From"`
	Host string  `toml:"Host"`
	User string  `toml:"User"`
	Password string  `toml:"Password"`
}

func NewMailHandler(config MailConfig) *MailHandler {
	mailHandler := MailHandler {
		from: config.From,
		host: config.Host,
		auth: smtp.PlainAuth("", config.User, config.Password, "smtp"),
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
		log.Print("faild send mail : %s", err.Error())
		return err
	}
	return nil
}
