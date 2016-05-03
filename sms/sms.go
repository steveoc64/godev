package sms

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/steveoc64/godev/config"
)

func GetBalance() (int, error) {
	c := config.Get()

	if c.SMSOn == false {
		return 0, nil
	}

	resp, err := http.PostForm(
		c.SMSServer,
		url.Values{
			"username": {c.SMSUser},
			"password": {c.SMSPasswd},
			"action":   {"balance"},
		})

	if err != nil {
		log.Println("HTTP Post Error", err.Error())
		return 0, err
	}

	/*	log.Println(resp)
		log.Println("status", resp.Status)
		log.Println("status code", resp.StatusCode)
		log.Println("proto", resp.Proto)
		log.Println("major", resp.ProtoMajor)
		log.Println("minor", resp.ProtoMinor)
		log.Println("header", resp.Header)
		log.Println("content length", resp.ContentLength)
		log.Println("transfer", resp.TransferEncoding)
		log.Println("trailer", resp.Trailer)
		log.Println("close", resp.Close)
		log.Println("req", resp.Request)
		log.Println("tls", resp.TLS)
		log.Println("body", resp.Body)
	*/

	// read the response
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	isok := string(body[:3])
	if isok != "OK:" {
		return 0, errors.New("Status Code is" + isok)
	}
	s := string(body[3:])
	return strconv.Atoi(s)
}

func Send(number string, message string, ref string) error {

	c := config.Get()
	resp, err := http.PostForm(
		c.SMSServer,
		url.Values{
			"username": {c.SMSUser},
			"password": {c.SMSPasswd},
			"to":       {number},
			"from":     {"SBS Intl"},
			"ref":      {ref},
			"message":  {message},
		})

	if err != nil {
		log.Println("HTTP Post Error", err.Error())
		return err
	}

	// read the response
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	lines := strings.Split(string(body), "\n")
	for _, v := range lines {
		p := strings.Split(v, ":")
		switch p[0] {
		case "OK":
			log.Println("SMS OK", p[1], "ref", p[2])
		case "BAD":
			log.Println("SMS BAD", p[1], "reason", p[2])
			return errors.New(p[2])
		case "ERROR":
			log.Println("SMS ERROR", p[1])
			return errors.New(p[1])
		}
	}
	return nil
}
