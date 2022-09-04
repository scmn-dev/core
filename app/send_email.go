package app

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendMail is an helper to send mail all over the project
func SendMail(name, email string, subject, bodyHTML string) error {
	from := mail.NewEmail("$PASSWORD_MANAGER_NAME", "EMAIL")
	to := mail.NewEmail(name, email)
	bodyText := ""
	message := mail.NewSingleEmail(from, subject, to, bodyText, bodyHTML)
	client := sendgrid.NewSendClient("SENDGRID_API_KEY")
	res, err := client.Send(message)

	if err != nil {
		log.Println(err)
		return err
	} else {
		log.Println(res.StatusCode)
	}

	return nil
}
