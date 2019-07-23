package library

import (
	"log"
	"net/smtp"	
)
import "gopkg.in/gomail.v2"

type MailLibrary struct{}

//FUNGSI CREATE
func (lib *MailLibrary) KirimEmail(body string) {
	from := "didintri196@gmail.com"
	pass := "segopecel"
	to := "didintri196@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Email Telah Dikirimkan")
}

//FUNGSI CREATE
func (lib *MailLibrary) GoEmail(body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "admin@serverku.online")
	m.SetHeader("To", "didintri196@gmail.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Serverku.online")
	m.SetBody("text/html", "Hello <b>"+body+"</b>!")
	//m.Attach("/home/Alex/lolcat.jpg")
	
	d := gomail.NewDialer("mail.serverku.online", 465, "admin@serverku.online", "segopecel12")
	
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}


