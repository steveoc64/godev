package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/steveoc64/godev/config"
	// "github.com/steveoc64/godev/mail"
	"github.com/steveoc64/godev/db"
	"github.com/steveoc64/godev/echocors"
	"github.com/steveoc64/godev/sms"
	"github.com/steveoc64/godev/smt"
	"log"
)

func main() {

	cpus := smt.Init()
	fmt.Printf("Yo Ho Ho, here we Go on %d CPU cores\n", cpus)

	cfg := config.LoadConfig()

	db.Init(cfg.DataSourceName)

	smsbal, smserr := sms.GetBalance()
	if smserr != nil {
		log.Fatal("Cannot retrieve SMS account info", smserr.Error())
	}
	log.Println("... Remaining SMS Balance =", smsbal)

	/*
		MailChannel := mail.InitMailer()
			m := mail.NewMail()
			m.SetHeader("To", "jack@sprat.com")
			m.SetHeader("Subject", "Loan Application")
			m.SetBody("text/html", "Awww Snap ! Your Loan Application has been denied  :(")
			MailChannel <- m
	*/
	e := echo.New()
	if cfg.Debug {
		e.SetDebug(true)
	}
	echocors.Init(e, cfg.Debug)

	// Start the web server
	if cfg.Debug {
		fmt.Printf("... Starting Web Server on port %d", cfg.WebPort)
	}
	e.Run(fmt.Sprintf(":%d", cfg.WebPort))
}
