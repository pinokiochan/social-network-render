package utils

import (
	"fmt"
	"os"

	"github.com/pinokiochan/social-network-render/internal/logger"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var (
	smtpHost = os.Getenv("SMTP_HOST") // Берем из системных переменных
	smtpPort = 587                    // Используем стандартный порт для STARTTLS
	smtpUser = os.Getenv("SMTP_USER") // Берем из системных переменных
	smtpPass = os.Getenv("SMTP_PASS") // Берем из системных переменных
)

// SendEmail отправляет письмо с опциональным вложением
func SendEmail(to, subject, body, attachmentPath string) error {
	// Проверка наличия всех обязательных переменных окружения
	if smtpHost == "" || smtpUser == "" || smtpPass == "" {
		logger.Log.WithFields(logrus.Fields{
			"host": smtpHost != "",
			"user": smtpUser != "",
		}).Error("Missing SMTP configuration")
		return fmt.Errorf("SMTP configuration is missing in environment variables")
	}

	// Создание нового письма
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Прикрепление файла, если путь указан
	if attachmentPath != "" {
		m.Attach(attachmentPath)
	}

	// Настройка и отправка письма
	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	if err := d.DialAndSend(m); err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
			"to":    to,
		}).Error("Failed to send email")
		return fmt.Errorf("Failed to send email: %v", err)
	}

	// Лог успешной отправки письма
	logger.Log.WithFields(logrus.Fields{
		"to":      to,
		"subject": subject,
	}).Info("Email sent successfully")

	return nil
}
