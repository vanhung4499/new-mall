package email

import (
	"gopkg.in/mail.v2"
	conf "new-mall/config"
)

// Sender represents an email sender
type Sender struct {
	SmtpHost      string `json:"smtp_host"`
	SmtpEmailFrom string `json:"smtp_email_from"`
	SmtpPass      string `json:"smtp_pass"`
}

// NewEmailSender creates a new EmailSender instance with configuration
func NewEmailSender() *Sender {
	eConfig := conf.Config.Email
	return &Sender{
		SmtpHost:      eConfig.SmtpHost,
		SmtpEmailFrom: eConfig.SmtpEmail,
		SmtpPass:      eConfig.SmtpPass,
	}
}

// Send sends an email with the specified data, recipient email, and subject
func (s *Sender) Send(data, emailTo, subject string) error {
	m := mail.NewMessage()
	m.SetHeader("From", s.SmtpEmailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", data)

	// Create a new mail dialer
	d := mail.NewDialer(s.SmtpHost, 465, s.SmtpEmailFrom, s.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Dial and send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
