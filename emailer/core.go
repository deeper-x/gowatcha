package emailer

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/zreq3b/gowatcha/settings"
)

// send sends email to recipient
func sendEmail(rcpt, msg string) (bool, error) {
	user := os.Getenv("USER")
	passwd := os.Getenv("PASSWD")

	auth := smtp.PlainAuth("", user, passwd, settings.SMTPSERVER)

	smtpURI := fmt.Sprintf("%v:%v", settings.SMTPSERVER, settings.SMTPPORT)
	err := smtp.SendMail(smtpURI, auth, user, []string{rcpt}, []byte(msg))
	if err != nil {
		log.Panic(err)
		return false, err
	}

	return true, nil
}

// Notify send email if last one has been sent more than a X ago
func Notify(rcpt, msg string, lastTS int64) (bool, error) {
	now := time.Now().Unix()
	delta := now - lastTS

	if delta <= settings.NOTIFINTVL {
		return false, errors.New("notification already sent")
	}

	_, err := sendEmail(rcpt, msg)
	if err != nil {
		return false, err
	}

	return true, nil
}
