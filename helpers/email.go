package helpers

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendOTPEmail(to string, token string) error {
	// t := template.New("./email/emailOTP.html")

	var err error
	t, err := template.ParseFiles("./email/emailotp.html")
	if err != nil {
		log.Println(err)
		return err
	}
	var data struct {
		Token string
	} = struct {
		Token string
	}{
		Token: token,
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
	var err error
	t, err := template.ParseFiles("./email/email_id_visit.html")
	if err != nil {
		log.Println(err)
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, visitID); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	// log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Email Konfirmasi Kunjungan")
	msg.SetBody("text/html", result)

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil

}

func SendVisitNotif(to string, guestName string, visitID int) error {
	var err error
	t, err := template.ParseFiles("./email/notif_visit.html")
	if err != nil {
		log.Println(err)
		return err
	}
	var data struct {
		GuestName string
		VisitID   int
	} = struct {
		GuestName string
		VisitID   int
	}{
		GuestName: guestName,
		VisitID:   visitID,
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	// redirectLink := os.Getenv("REDIRECT_LINK")
	// log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Notifikasi Kunjungan")
	msg.SetBody("text/html", result)

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil

}

func SendCancelProposalEmail(to string, visitID int) error {
	var err error
	t, err := template.ParseFiles("./email/email_dec.html")
	if err != nil {
		log.Println(err)
		return err
	}
	var data struct {
		VisitID int
	} = struct {
		VisitID int
	}{
		VisitID: visitID,
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	fmt.Println("Cancel Proposal to : ", to)
	senderEmail := os.Getenv("SERVICE_EMAIL")
	senderPass := os.Getenv("SERVICE_EMAIL_PASS")
	// log.Printf("Email:%s\nPassword:%s\n", senderEmail, senderPass)
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderEmail)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Kunjungan Ditolak")
	msg.SetBody("text/html", result)

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}

func SendConfirmProposalEmail(to string, visitID int) error {
	var err error
	t, err := template.ParseFiles("./email/email_acc.html")
	if err != nil {
		log.Println(err)
		return err
	}
	var data struct {
		VisitID int
	} = struct {
		VisitID int
	}{
		VisitID: visitID,
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
	msg.SetHeader("Subject", "Kunjungan Diterima")
	msg.SetBody("text/html", result)

	n := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPass)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
