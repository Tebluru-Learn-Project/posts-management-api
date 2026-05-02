package service

import (
	"fmt"
	"net/smtp"
	"go-api/config"
)

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func (s *MailService) SendOTP(to string, code string) error {
	host := config.GetEnv("SMTP_HOST", "")
	port := config.GetEnv("SMTP_PORT", "587")
	user := config.GetEnv("SMTP_USER", "")
	pass := config.GetEnv("SMTP_PASS", "")
	from := config.GetEnv("SMTP_FROM", user)

	auth := smtp.PlainAuth("", user, pass, host)

	subject := "Subject: Verify Your Account\r\n"
	body := fmt.Sprintf("Your OTP code is: %s\nThis code will expire in 5 minutes.", code)

	message := []byte(subject + "\r\n" + body)

	return smtp.SendMail(host+":"+port, auth, from, []string{to}, message)
}