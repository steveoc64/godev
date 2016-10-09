package mail

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/steveoc64/godev/config"
	gomail "gopkg.in/gomail.v2"
)

var MailChannel = make(chan *gomail.Message, 64)

func NewMail() *gomail.Message {

	m := gomail.NewMessage()
	m.SetHeader("From", "umpire@wargaming.io")
	return m
}

func InitMailer() chan *gomail.Message {
	go mailerDaemon()
	return MailChannel
}

func mailerDaemon() {

	c := config.Get()
	log.Println("starting mailer with", c.MailServer, c.MailPort, c.MailUser, c.MailPasswd)
	d := gomail.NewPlainDialer(
		c.MailServer,
		c.MailPort,
		c.MailUser,
		c.MailPasswd)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	var s gomail.SendCloser
	var err error
	open := false
	for {
		select {
		case m, ok := <-MailChannel:
			if !ok {
				return
			}
			if !open {
				if s, err = d.Dial(); err != nil {
					panic(err)
				}
				open = true
			}
			if err := gomail.Send(s, m); err != nil {
				log.Print(err)
			}
		// Close the connection to the SMTP server if no email was sent in
		// the last couple of minutes.
		case <-time.After(120 * time.Second):
			if open {
				if err := s.Close(); err != nil {
					panic(err)
				}
				open = false
			}
		}
	}
}
