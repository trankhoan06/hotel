package email

import (
	"fmt"
)

func SendCode(email string, code int) error {
	title := "VERIFY EMAIL"
	content := fmt.Sprintf("email confirmation code %d", code)
	to :=
		[]string{
			email,
		}
	sender := NewGmailSender()
	if err := sender.SendEmail(title, content, to, nil, nil, nil); err != nil {
		return err
	}
	return nil
}
