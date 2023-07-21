package email

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
)

// Email struct used to organize our email service
type Email struct {
	From     string
	Password string
	Host     string
	Port     string
}

func SendSignupVerificationEmail(email Email, to []string, name string, link string) {
	// Authentication
	auth := smtp.PlainAuth("", email.From, email.Password, email.Host)

	// Note: templates directory is within cmd/ directory
	t, err := template.ParseFiles("templates/sign_up_verification_email.html")
	if err != nil {
		log.Println("Template does not exist")
		return
	}
	items := struct {
		Name string
		URL  string
	}{
		Name: name,
		URL:  link,
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, items); err != nil {
		log.Println(err)
	}
	body := buf.String()
	subject := "Subject: Confirm your account!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + body)

	// Sending email.

	emailErr := smtp.SendMail(email.Host+":"+email.Port, auth, email.From, to, msg)
	if emailErr != nil {
		log.Println(emailErr)
		return
	}
	log.Println("Email successfully sent")
}
