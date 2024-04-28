package user_gateways

import (
	"errors"
	"os"

	log_messages "github.com/glener10/authentication/src/log/messages"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
	gomail "gopkg.in/mail.v2"
)

func SendEmail(to string, subject string, code string) error {
	mail := gomail.NewMessage()

	mail.SetHeader("From", os.Getenv("EMAIL_FROM"))
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", "your code is: "+code)

	sendEmail := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_PASSWORD"))
	sendEmail.StartTLSPolicy = gomail.MandatoryStartTLS

	if err := sendEmail.DialAndSend(mail); err != nil {
		go utils_usecases.CreateLog(nil, "sendemail", "POST", false, log_messages.SEND_EMAIL_WITHOUT_SUCCESS, "")
		return errors.New("error to send email")
	}

	return nil
}
