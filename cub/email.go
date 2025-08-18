package main

import (
	"fmt"
	"net/smtp"
)

func ListenForEmail() {
	for {
		select {}
	}
}

func SendEmail(to, subject, body string, error_channel chan error) {
	msg := fmt.Sprintf("From: %s\r\n", "support@dailywire.com") +
		fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		body
	err := smtp.SendMail("localhost:1025", nil, "support@dailywire.com", []string{to}, []byte(msg))
	if err != nil {
		error_channel <- err
		return
	}
}
