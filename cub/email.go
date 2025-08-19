package main

import (
	"fmt"
	"net/smtp"
	"sync"
)

func ListenForEmailErrors(email_error_channel chan error) {
	for {
		err := <-email_error_channel
		fmt.Printf("Error sending email: %v\n", err)
	}
}

func SendEmail(to, subject, body string, error_channel chan error, wg *sync.WaitGroup) {
	defer wg.Done()
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

func (app *Bridge) SendEmail(to, subject, body string) {
	app.WaitGroup.Add(1)
	go SendEmail(to, subject, body, app.EmailErrChannel, app.WaitGroup)
}
