package email

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct{}

func (sg SendGrid) SendSimple(to string, subject string, body string) error {
	from := mail.NewEmail("Example User", "test@example.com")

	receipient := mail.NewEmail("Example User", to)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, receipient, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	if _, err := client.Send(message); err != nil {
		return err
	} else {
		return nil
	}

}
