package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
)

func main() {
	const required = "REQUIRED"

	apiUser := flag.String("apiUser", required, "Username to connect to the SendGrid API")
	apiKey := flag.String("apiKey", required, "Password to connect to the SendGrid API")
	fromEmail := flag.String("fromEmail", required, "Sender of the test email")
	toEmail := flag.String("toEmail", required, "Recipient of the test email")
	flag.Parse()

	if *apiUser == required || *apiKey == required || *fromEmail == required || *toEmail == required {
		fmt.Println("apiUser, apiKey, fromEmail and toEmail are required")
		return
	}

	// Set up authentication information.
	log.Println("Creating authentication")
	auth := smtp.PlainAuth(
		"",
		*apiUser,
		*apiKey,
		"smtp.sendgrid.net",
	)

	log.Println("Creating email")
	fromName := strings.Split(*fromEmail, "@")[0]
	toName := strings.Split(*toEmail, "@")[0]

	from := mail.Address{Name: fromName, Address: *fromEmail}
	to := mail.Address{Name: toName, Address: *toEmail}

	title := "SendGrid SMTP Test"
	body := "This email has been sent via SMTP through SendGrid"

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	log.Printf("Sending email to %s from %s\n", to.Address, from.Address)
	err := smtp.SendMail(
		"smtp.sendgrid.net:587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Finished")
}
