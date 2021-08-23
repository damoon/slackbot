package bot

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/slack-go/slack"
)

type email struct{}

func (s email) Response(msg *slack.MessageEvent) string {
	if !strings.Contains(strings.ToLower(msg.Text), "email") {
		return ""
	}

	message, err := sendEmail()
	if err != nil {
		log.Printf("sending email: %v", err)
		return "message was forwarded to email alerting"
	}

	return message
}

func sendEmail() (string, error) {
	// TODO: needs config

	// Sender data.
	from := "from@gmail.com"
	password := "<Email Password>"

	// Receiver email address.
	to := []string{
		"sender@example.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return "", err
	}
	fmt.Println("Email Sent Successfully!")

	return "", nil
}
