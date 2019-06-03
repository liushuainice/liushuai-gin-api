package mail

import (
	"github.com/go-gomail/gomail"
)

// SMTPConfig 用于解析smtp配置
type SMTPConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Account string `json:"account"`
	Captcha string `json:"captcha"`
}

// SMTP Use dialer send a mail
var SMTP *gomail.Dialer

// InitSMTPCli create a dialer handler for martini
func InitSMTPCli(conf *SMTPConfig) {
	SMTP = gomail.NewPlainDialer(conf.Address, conf.Port, conf.Account, conf.Captcha)
}

// SendMail 发送邮件
func SendMail(messgae *gomail.Message) error {
	// 拨号并发送邮件
	err := SMTP.DialAndSend(messgae)
	if err != nil {
		println(err)
		return err
	}
	return nil
}
