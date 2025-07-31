package main

import (
	"io"
	"net/mail"
	"strings"
)

func ExtractMailContent(mailData string) string {
	reader := strings.NewReader(mailData)
	msg, err := mail.ReadMessage(reader)
	if err != nil {
		return ""
	}
	bytes, err := io.ReadAll(msg.Body)
	if err != nil {
		return ""
	}
	return string(bytes)
}
