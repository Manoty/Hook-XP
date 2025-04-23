package infra



type  OTPSender interface {
	SendOTP(email, code string) error
}