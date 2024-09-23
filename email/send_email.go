package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	NameSender          = "Hotel"
	SenderEmailAccount  = "trankhoan06@gmail.com"
	PassWordSenderEmail = "klptazdfrrcwvgfp"
	SmtpAuthorAddress   = "smtp.gmail.com"
	SmtpSeverService    = "smtp.gmail.com:587"
)

type Sender interface {
	SendEmail(Title string,
		Content string,
		to []string,
		Cc []string,
		Bcc []string,
		AttactFile []string) error
}
type GmailSender struct {
	Name              string
	FromEmail         string
	FromEmailPassword string
}

func NewGmailSender() Sender {
	return &GmailSender{
		Name:              NameSender,
		FromEmail:         SenderEmailAccount,
		FromEmailPassword: PassWordSenderEmail,
	}
}

func (sender *GmailSender) SendEmail(Title string,
	Content string,
	to []string,
	Cc []string,
	Bcc []string,
	AttactFile []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", NameSender, SenderEmailAccount)
	e.To = to
	e.Subject = Title
	e.Text = []byte(Content)
	e.Bcc = Bcc
	for _, f := range AttactFile {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s:%w", f, err)
		}

	}
	authau := smtp.PlainAuth("", sender.FromEmail, sender.FromEmailPassword, SmtpAuthorAddress)
	return e.Send(SmtpSeverService, authau)
}
