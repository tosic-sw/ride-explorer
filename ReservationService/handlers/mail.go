package handlers

import (
	"ReservationService/models"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

type MailHandler struct {
	host     string
	port     int
	address  string
	from     string
	password string
}

func NewMailHandler() *MailHandler {
	return &MailHandler{
		host:     "smtp-mail.outlook.com",
		port:     587,
		from:     "imejlposlanic@outlook.com",
		password: "MailNuget123",
	}
}

func (mh *MailHandler) SendMail(mail models.Mail) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", mh.from)

	// Set E-Mail receivers
	m.SetHeader("To", mail.To)

	// Set E-Mail subject
	m.SetHeader("Subject", mail.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", mail.Body)

	// Settings for SMTP server
	d := gomail.NewDialer(mh.host, mh.port, mh.from, mh.password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
