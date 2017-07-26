package main

import (
	"flag"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

func checkHost(email string) string {

	host := strings.Split(email, "@")
	mx, err := net.LookupMX(host[1])
	if err != nil {
		return `Failed`
	}

	client, err := smtp.Dial(fmt.Sprintf("%s:%d", mx[0].Host, 25))
	defer client.Close()
	if err != nil {
		return `Failed`
	}
	err = client.Mail(email)
	if err != nil {
		return `Failed`
	}

	return `OK`
}

func main() {
	var emailAddy string
	flag.StringVar(&emailAddy, `email`, `null@null`, `Please enter an email address`)
	flag.Parse()
	fmt.Println(checkHost(emailAddy))
}
