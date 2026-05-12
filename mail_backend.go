package main

import (
	"io"
	"slices"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type MailBackend struct {
	Domain  string
	Storage Storage
}

func (bkd *MailBackend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &MailSession{Domain: bkd.Domain, Storage: bkd.Storage, Conn: c}, nil
}

type MailSession struct {
	Domain  string
	Storage Storage
	Conn    *smtp.Conn
	mail    string
}

func (s *MailSession) AuthMechanisms() []string {
	return []string{sasl.Plain}
}

// Auth is the handler for supported authenticators.
func (s *MailSession) Auth(mech string) (sasl.Server, error) {
	return sasl.NewPlainServer(func(identity, username, password string) error {
		return nil
	}), nil
}

func (s *MailSession) Mail(from string, opts *smtp.MailOptions) error {
	return nil
}

func (s *MailSession) Rcpt(to string, opts *smtp.RcptOptions) error {
	parts := strings.Split(to, "@")
	if len(parts) != 2 {
		defer s.Conn.Close()
		return smtp.ErrServerClosed
	}
	if parts[1] != s.Domain {
		defer s.Conn.Close()
		return smtp.ErrServerClosed
	}
	mails, _ := s.Storage.GetMails()
	if !slices.Contains(mails, parts[0]) {
		defer s.Conn.Close()
		return smtp.ErrServerClosed
	}
	s.mail = parts[0]
	return nil
}

func (s *MailSession) Data(r io.Reader) error {
	if b, err := io.ReadAll(r); err != nil {
		return err
	} else {
		s.Storage.StoreMailContent(s.mail, ExtractMailContent(string(b)))
	}
	return nil
}

func (s *MailSession) Reset() {}

func (s *MailSession) Logout() error {
	return nil
}
