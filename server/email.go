package main

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

func sendEmail(setting EmailSetting, recipient, subject, body string) error {
	addr := fmt.Sprintf("%s:%d", setting.SMTPHost, setting.SMTPPort)
	var auth smtp.Auth
	if setting.SMTPUser != "" {
		auth = smtp.PlainAuth("", setting.SMTPUser, setting.SMTPPassword, setting.SMTPHost)
	}
	from := setting.SenderEmail
	message := strings.Join([]string{
		"From: " + from,
		"To: " + recipient,
		"Subject: " + subject,
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"",
		body,
	}, "\r\n")
	if !setting.IsSSL {
		return smtp.SendMail(addr, auth, from, []string{recipient}, []byte(message))
	}

	conn, err := tls.Dial("tcp", addr, &tls.Config{ServerName: setting.SMTPHost, MinVersion: tls.VersionTLS12})
	if err != nil {
		return err
	}
	defer conn.Close()
	client, err := smtp.NewClient(conn, setting.SMTPHost)
	if err != nil {
		return err
	}
	defer client.Quit()
	if setting.SMTPUser != "" {
		if err := client.Auth(auth); err != nil {
			return err
		}
	}
	if err := client.Mail(from); err != nil {
		return err
	}
	if err := client.Rcpt(recipient); err != nil {
		return err
	}
	writer, err := client.Data()
	if err != nil {
		return err
	}
	if _, err := writer.Write([]byte(message)); err != nil {
		_ = writer.Close()
		return err
	}
	return writer.Close()
}
