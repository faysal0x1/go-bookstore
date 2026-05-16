package services

import (
	"log"
	"strconv"

	"github.com/faysal0x1/go-bookstore/pkg/config"
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendWelcomeEmail(toEmail, userName string) error
}

type emailService struct {
	config *config.Config
}

func NewEmailService(config *config.Config) EmailService {
	return &emailService{config: config}
}

func (s *emailService) SendWelcomeEmail(toEmail, userName string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.config.SMTPSender)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Welcome to Bookstore!")
	m.SetBody("text/html", "<h1>Hello "+userName+"</h1><p>Thanks for joining our Bookstore!</p>")

	port, _ := strconv.Atoi(s.config.SMTPPort)
	d := gomail.NewDialer(s.config.SMTPHost, port, s.config.SMTPUser, s.config.SMTPPass)

	// Send the email (In a real app, you might do this in a goroutine)
	go func() {
		if err := d.DialAndSend(m); err != nil {
			log.Printf("Failed to send welcome email to %s: %v", toEmail, err)
		} else {
			log.Printf("Welcome email sent to %s", toEmail)
		}
	}()

	return nil
}
