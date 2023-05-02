package email

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var (
	EMAIL_SMTP_HOST = os.Getenv("EMAIL_SMTP_HOST")
	EMAIL_SMTP_PORT = os.Getenv("EMAIL_SMTP_PORT")
	EMAIL_USER      = os.Getenv("EMAIL_USER")
	EMAIL_PWD       = os.Getenv("EMAIL_PWD")
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "1481700559@qq.com")
	m.SetHeader("To", "magicminorn@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>小N</b>!")
	var d *gomail.Dialer
	if port, err := strconv.Atoi(EMAIL_SMTP_PORT); err != nil {
		log.Fatalln("EMAIL_SMTP_PORT is not a number")
	} else {
		d = gomail.NewDialer(EMAIL_SMTP_HOST, port, EMAIL_USER, EMAIL_PWD)
	}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
