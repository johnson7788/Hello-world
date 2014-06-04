package main

import (
	"net/smtp"
	"log"
)

func main() {
	auth := smtp.PlainAuth(
		"",
		"monitorlog@126.com",
		"password",
		"smtp.126.com",
	)

	err := smtp.SendMail(
		"smtp.126.com:25",
		auth,
		"monitorlog@126.com",
		[]string{"monitorlog@126.com"},
		[]byte("This is the email"),
	)
	if err !=nil {
		log.Fatal(err)
	}
}


