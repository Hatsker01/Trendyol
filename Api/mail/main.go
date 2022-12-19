package mail

import (
	"net/smtp"
)

func SendMail(code string, email string) error {

	from := "jamshidbek1805@gmail.com"
	password := "qzsxlgudmntzsqhs"

	toEmailAddress := email
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Your verification code is: \n"
	body := code
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil

}
