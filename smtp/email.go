package email

import (
	"log"
	"net/smtp"
	"time"
)

// func email() {
// 	Send("hello there")
// }

// Sends an email
func Send(body string) {
	from := "test@gmail.com"
	pass := "pass"
	to := "test2@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %q", err)
		return
	}

	log.Printf("Mail sent to %s, from %s, on %q", []string{to}, []string{from}, time.Now())
}
