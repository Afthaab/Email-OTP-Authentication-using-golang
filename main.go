package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	gomail "gopkg.in/mail.v2"
)

func genCaptchaCode() (string, error) {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}

func main() {

	//initializing env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get the email from the user :
	fmt.Println("Enter The Email Address : ")
	var Emmail string
	fmt.Scan(&Emmail)
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", Emmail)

	// Set E-Mail receivers
	m.SetHeader("To", os.Getenv("RECIEVEREMAIL"))

	// Set E-Mail subject
	m.SetHeader("Subject", "OTP to verify your Gmail")

	//otp generation
	onetimepassword, err := genCaptchaCode()
	if err != nil {
		fmt.Println(err)
	}

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", onetimepassword+" is your OTP to register to our site. Thank you registering to our site. Happy Shopping :)")

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
