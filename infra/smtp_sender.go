package infra

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// SmtpSender is a struct that implements the OTPSender interface
type SmtpSender struct {
	Host   string
	Port  int
	Username string
	Password string
	From string
}

// sendOTP sends an OTP to the given email address using SMTP
func (s *SmtpSender) SendOTP(email, code string) error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	// create the email message
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: Your OTP Code\r\n\r\nYour OTP code is: %s", email, code))
	
	// send the email
	err := smtp.SendMail(addr, auth, s.From, []string{email}, msg)
	if err != nil {
		log.Printf("send otp failed for  %s: %v", email, err)
		return err
	}
	return nil

}
//NewOTPSender initializes and returns a configured smtp sender
func NewOTPSender() *SmtpSender {
	return &SmtpSender{
	Host : "smtp.gmail.com",
	Port : 587,
	Username : "info.freshlyfarms5@gmail.com",
	//Password : "jbzh yrga gils yabk",
	Password : os.Getenv("SMTP_PASSWORD"),
	From: "info.freshlyfarms5@gmail.com",
	}


}
