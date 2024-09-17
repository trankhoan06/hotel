package email

type Sender interface {
	SendEmail(Title string,
		Content string,
		to []string,
		Cc []string,
		Bcc []string,
		AttactFile []string) error
}
