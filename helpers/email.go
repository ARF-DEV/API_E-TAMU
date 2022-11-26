package helpers

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(to string, data interface{}) error {
	// t := template.New("./email/emailOTP.html")

	var err error
	t, err := template.ParseFiles("./email/emailotp.html")
	if err != nil {
		log.Println(err)
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	// log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Nomor OTP Pengajuan Kunjungan")
	msg.SetBody("text/html", result)

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func SendVisitID(to string, visitID int) error {
	// fmt.Println("Visit ID to : ", to)
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	// log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Email Konfirmasi Kunjungan")
	msg.SetBody("text/html", fmt.Sprintf("ID Kunjungan: <b>%d</b>", visitID))

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil

}

func SendVisitNotif(to string, guestName string, visitID int) error {
	// fmt.Println("Notif to : ", to)
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	// redirectLink := os.Getenv("REDIRECT_LINK")
	// log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Notifikasi Kunjungan")
	msg.SetBody("text/html", fmt.Sprintf(`Tamu atas nama <b>%s</b> ingin bertemu dengan anda </br>
	klik <a href='http://localhost:8000/api/v1/visits/confirmvisit?id=%d'>link ini</a> untuk menerima permintaan ini</br>
	klik <a href='http://localhost:8000/api/v1/visits/cancelvisit?id=%d'>link ini </a> untuk menolak permintaan ini`, guestName, visitID, visitID))

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil

}

func SendCancelProposalEmail(to string, visitID int) error {
	fmt.Println("Cancel Proposal to : ", to)
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	// log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Kunjungan Ditolak")
	msg.SetBody("text/html", fmt.Sprintf("Kunjungan anda dengan id <b>%d</b>, telah ditolak oleh yang dikunjungi", visitID))

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}

func SendConfirmProposalEmail(to string, visitID int) error {
	// fmt.Println("Visit ID to : ", to)
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	// log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Kunjungan Diterima")
	msg.SetBody("text/html", fmt.Sprintf("Kunjungan anda dengan id <b>%d</b>, telah diterima oleh yang dikunjungi", visitID))

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
