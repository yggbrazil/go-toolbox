package email

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"path/filepath"
	"strconv"
	"strings"
)

type Service interface {
	SendEmail(m Mail) error
}

type service struct {
	SmtpHost string
	SmtpPort int
	From     string
	Password string
}

type Mail struct {
	Subject     string
	To          []string
	CC          []string
	BCC         []string
	Body        string
	Attachments []string
}

func New(smtpHost string, smtpPort int, from, password string) (*service, error) {
	return &service{
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
		From:     from,
		Password: password,
	}, nil
}

func (s service) SendEmail(m Mail) error {
	addr := s.SmtpHost + ":" + strconv.Itoa(s.SmtpPort)
	auth := smtp.PlainAuth("", s.From, s.Password, s.SmtpHost)
	message, err := BuildMessage(m)
	if err != nil {
		return err
	}

	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         s.SmtpHost,
	}

	c, err := smtp.Dial(addr)
	if err != nil {
		log.Panic(err)
	}

	c.StartTLS(tlsconfig)

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(s.From); err != nil {
		return err
	}

	if err = c.Rcpt(strings.Join(m.To, ";")); err != nil {
		return err
	}

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

	c.Quit()

	return nil
}

func BuildMessage(m Mail) (string, error) {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()

	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	}

	buf.WriteString("Content-Type: text/html; charset=utf-8\n\n")
	buf.WriteString(m.Body)

	if withAttachments {
		for _, attachment := range m.Attachments {
			fileName, fileData, err := readFile(attachment)
			if err != nil {
				return "", err
			}

			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(fileData)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n\n", fileName))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(fileData)))
			base64.StdEncoding.Encode(b, fileData)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	return buf.String(), nil
}

func readFile(filePath string) (string, []byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", []byte{}, fmt.Errorf("error on find file: %s", filePath)
	}

	_, fileName := filepath.Split(filePath)

	return fileName, data, nil
}
