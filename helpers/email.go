package helpers

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(to string, otpString string) error {
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "WoiShop Registeration OTP")
	msg.SetBody("text/html", fmt.Sprintf("Code : <b>%s</b>", otpString))

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil

}
