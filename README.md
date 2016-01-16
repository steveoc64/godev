# godev
Development utils for GO 

Tiny but useful code snippets to link into your GO project.

Import / Fork / Cut-n-Paste .... up to you.

```go
package main

import (
	"fmt"
	"github.com/steveoc64/godev/config"
	"github.com/steveoc64/godev/mail"
)

func main() {
	...
```


## godev/config

Use a config.json file to store your apps operational params, and load them at runtime.

```go
func main() {

	cfg := config.LoadConfig()
	...
```

Loads the following file
```json
{
	"Debug": false,
	"DataSourceName": "user=postgres password=xxxxxx dbname=dbdbdb sslmode=disable",
	"WebPort": 8000,
	"MailServer": "mail.mysite.com",
	"MailUser": "mailaccount",
	"MailPasswd": "xxxxxx",
	"MailPort": 465
}
```

## godev/mail

Extend gopkg.in/gomail.v2 by adding an internal mailserver that uses channels.
Runtime config requires the config package.

```go
package main

import (
	"github.com/steveoc64/godev/config"
	"github.com/steveoc64/godev/mail"
)

func main() {

	cfg := config.LoadConfig()
	MailChannel := mail.InitMailer()

	m := mail.NewMail()
	m.SetHeader("To", "jack@sprat.com")
	m.SetHeader("Subject", "Loan Application")
	m.SetBody("text/html", "Awww Snap ! Your Loan Application has been denied  :(")
	MailChannel <- m
```

So inside your app, you create emails using the regular gomail funtions, and then just add them to the MailChannel 
queue.  A goroutine in the background works through the MailChannel queue, and does all the messy delivery stuff 
for you.

## SMT

## JWT

## Cors

## DB
