# godev
Development utils for GO 

Tiny Libs to link into your GO project


## Config

Use a config.json file to store your apps operational params, and load them at runtime.

```go
func main() {

	cfg := config.LoadConfig()
	...
```

## Mail

Extend gopkg.in/gomail.v2 by adding an internal mailserver that uses channels.
Runtime config requires the config package.

## SMT

## JWT

## Cors

## DB
