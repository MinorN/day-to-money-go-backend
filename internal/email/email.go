package email

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func newDialer() *gomail.Dialer {
	return gomail.NewDialer(
		viper.GetString("email.smtp.host"),
		viper.GetInt("email.smtp.port"),
		viper.GetString("email.smtp.user"),
		viper.GetString("email.smtp.password"),
	)
}

func newMessage(to, subject, body string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", "1481700559@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	return m
}

func Send() {
	m := newMessage("magicminorn@gmail.com", "Hello!", "Hello <b>小N</b>!")

	d := newDialer()
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func SendValidationCode(email, code string) error {
	m := newMessage(email, fmt.Sprintf("[%s] 山竹记账验证码", code),
		fmt.Sprintf(`
		你正在登录或者注册山竹记账，你的验证码是 %s 。
		<br />
		如果你没有进行相关操作，请直接忽略本邮件即可`, code),
	)

	d := newDialer()
	return d.DialAndSend(m)
}
