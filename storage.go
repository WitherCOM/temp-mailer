package main

import (
	"errors"
	"maps"
	"slices"
)

type Storage interface {
	GetMails() ([]string, error)
	AssignMail(mail string) error
	StoreMailContent(mail string, content string) error
	GetMailContent(mail string) (string, error)
}

type InMemoryStorage struct {
	mails map[string]string
}

func InitInMemoryStorage() *InMemoryStorage {
	storage := &InMemoryStorage{}
	storage.mails = make(map[string]string)
	return storage
}

func (s *InMemoryStorage) GetMails() ([]string, error) {
	return slices.Collect(maps.Keys(s.mails)), nil
}

func (s *InMemoryStorage) AssignMail(mail string) error {
	s.mails[mail] = ""
	return nil
}

func (s *InMemoryStorage) StoreMailContent(mail string, content string) error {
	s.mails[mail] = content
	return nil
}

func (s *InMemoryStorage) GetMailContent(mail string) (string, error) {
	if content, ok := s.mails[mail]; ok {
		return content, nil
	}
	return "", errors.New("mail not found")
}
