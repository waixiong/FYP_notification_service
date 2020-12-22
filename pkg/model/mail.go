package model

import (

	//pb "./proto"

	// "github.com/syndtr/goleveldb/leveldb"

	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

// SendNoReplyMail - Send one email, with html body
func (m *MailModel) SendNoReplyMail(ctx context.Context, repName string, repAddress string, subj string, body string) error {
	from := mail.Address{Name: "GetIt", Address: "wx.chee@getitqec.com"}
	to := mail.Address{Name: repName, Address: repAddress}
	// subj := "This is the email subject"
	// body := "This is an example body.\n With three lines.\nThis email is auth with md5, safer."

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html"
	headers["charset"] = "UTF-8"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "smtp.zoho.com:587"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", from.Address, "q17QrIjrYsAZ", host)

	// TLS config
	tlsconfig := &tls.Config{
		// InsecureSkipVerify: true, // perform checking for security
		ServerName: host,
	}

	c, err := smtp.Dial(servername)
	if err != nil {
		return err
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
