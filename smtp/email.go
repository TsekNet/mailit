package email

import (
	"log"
	"net/smtp"
	"time"
)

// Send an email
func Send(body string) {
	from := "email@email.com"
	pass := "pass"
	to := "email@email.com"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	server := "smtp.gmail.com:587"
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	subject := "Subject: Test email from Go!\n"

	msg := []byte(subject + mime + "<html><body><h1>" + body + "</h1></body></html>")

	err := smtp.SendMail(server, auth, from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %q", err)
		return
	}

	log.Printf("Mail sent to %s, from %s, on %q", []string{to}, []string{from}, time.Now())
}
