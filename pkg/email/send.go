package email

import "gopkg.in/gomail.v2"

const (
	host     = "smtp.exmail.qq.com"
	port     = 465
	username = "username"
	password = "password"
	sender   = "admin@example.com"
)

func Send(subject, content string, address ...string) error {
	gm := gomail.NewMessage()
	gm.SetHeader("From", sender)
	gm.SetHeader("To", address...)
	gm.SetHeader("Subject", subject)
	gm.SetBody("text/html", content)
	d := gomail.NewDialer(host, port, username, password)
	return d.DialAndSend(gm)
}
