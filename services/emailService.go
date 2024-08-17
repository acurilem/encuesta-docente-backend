package services

import (
	"log"
	"os"
	"strconv"

	"github.com/wneessen/go-mail"
)

func SendEmail(to string, subject string, body string) bool {
	m := mail.NewMsg()
	if err := m.From(os.Getenv("MAIL_FROM")); err != nil {
		log.Fatalf("failed to set From address: %s", err)
		return false
	}
	if err := m.To(to); err != nil {
		log.Fatalf("failed to set To address: %s", err)
		return false
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextHTML, body)
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	c, err := mail.NewClient(os.Getenv("MAIL_HOST"), mail.WithPort(port), mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithUsername(os.Getenv("MAIL_USERNAME")), mail.WithPassword(os.Getenv("MAIL_PASSWORD")))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
		return false
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
		return false
	}
	return true
}
