package email

type MailSender interface {
	SendSimple(to string, subject string, body string) error
	SendHTML(to []string, subject string, data map[string]interface{}, template string) error
}
