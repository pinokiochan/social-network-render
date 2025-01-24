package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/joho/godotenv"
	"github.com/pinokiochan/social-network-render/internal/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func SendEmail(to, subject, body, attachmentPath string) error {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to load .env file")
		return fmt.Errorf("Error loading .env file: %v", err)
	}

	// Извлечение SMTP настроек из окружения
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := 587 // Используем STARTTLS
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	// Проверка наличия всех обязательных переменных окружения
	if smtpHost == "" || smtpUser == "" || smtpPass == "" {
		logger.Log.WithFields(logrus.Fields{
			"host": smtpHost != "",
			"port": smtpPort,
			"user": smtpUser != "",
		}).Error("Missing SMTP configuration")
		return fmt.Errorf("SMTP configuration is missing in environment variables")
	}

	// Создание нового письма
	e := email.NewEmail()
	e.From = smtpUser
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)

	// Прикрепление файла, если путь указан
	if attachmentPath != "" {
		_, err := e.AttachFile(attachmentPath)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"error": err.Error(),
				"path":  attachmentPath,
			}).Error("Failed to attach file to email")
			return fmt.Errorf("failed to attach file: %v", err)
		}
	}

	// Установка TLS-соединения
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	address := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	err = e.SendWithTLS(address, smtp.PlainAuth("", smtpUser, smtpPass, smtpHost), tlsConfig)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
			"to":    to,
		}).Error("Failed to send email")
		return fmt.Errorf("failed to send email: %v", err)
	}

	// Логирование успешной отправки письма
	logger.Log.WithFields(logrus.Fields{
		"to":      to,
		"subject": subject,
	}).Info("Email sent successfully")

	return nil
}
