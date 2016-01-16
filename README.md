# godev
Development utils for GO 

Tiny but useful code snippets to link into your GO project

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


## github.com/steveoc64/godev/config

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

## github.com/steveoc64/godev/mail

Extend gopkg.in/gomail.v2 by adding an internal mailserver that uses channels.
Runtime config requires the config package.

```go
```


## SMT

## JWT

## Cors

## DB
