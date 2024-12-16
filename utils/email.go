package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-mail/mail"
)

// GenerateOTP generates a 6-digit OTP
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// SendOTPEmail sends the OTP email
func SendOTPEmail(email string, otp string) error {
	// Get SMTP configuration from env
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	// Validate env variables
	if smtpHost == "" || smtpPort == "" || smtpEmail == "" || smtpPassword == "" {
		return fmt.Errorf("missing SMTP configuration in environment variables")
	}

	// Convert port to integer
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT value: %v", err)
	}

	// Setup SMTP dialer
	dialer := mail.NewDialer(smtpHost, port, smtpEmail, smtpPassword)

	// Create the email message
	msg := mail.NewMessage()
	msg.SetHeader("From", smtpEmail)
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Your OTP for Password Reset")
	msg.SetBody("text/plain", fmt.Sprintf("Your OTP is: %s", otp))

	// Send the email
	if err := dialer.DialAndSend(msg); err != nil {
		log.Printf("Error sending OTP email to %s: %v", email, err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
