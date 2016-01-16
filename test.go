package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/steveoc64/godev/config"
	"github.com/steveoc64/godev/mail"
)

func main() {

	cfg := config.LoadConfig()
	mail.InitMailer()

	e := echo.New()
	if cfg.Debug {
		e.SetDebug(true)
	}

	// Start the web server
	if cfg.Debug {
		fmt.Printf("... Starting Web Server on port %d", cfg.WebPort)
	}
	e.Run(fmt.Sprintf(":%d", cfg.WebPort))
}
