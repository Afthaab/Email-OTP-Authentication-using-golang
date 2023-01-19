package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	gomail "gopkg.in/mail.v2"
)

func main() {

	//initializing env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", os.Getenv("SENDEREMAIL"))

	// Set E-Mail receivers
	m.SetHeader("To", os.Getenv("RECIEVEREMAIL"))

	// Set E-Mail subject
	m.SetHeader("Subject", "OTP to verify your Gmail")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "741741 is your OTP to register to our site. Thank you registering to our site. Happy Shopping :)")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SENDEREMAIL"), os.Getenv("PASSWORD"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("OTP has been sent successfully")
	}
}
