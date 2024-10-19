package email

import (
	"fmt"
)

func SendCodeForgot(email string, code int) error {
	title := "FORGOT PASSWORD"
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
