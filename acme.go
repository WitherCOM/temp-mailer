package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
)

type AcmeFile struct {
	LetsEncrypt Resolver `json:"letsencrypt"`
}

type Resolver struct {
	Certificates []CertEntry `json:"Certificates"`
}

type CertEntry struct {
	Domain struct {
		Main string `json:"main"`
	} `json:"domain"`

	Certificate string `json:"certificate"`
	Key         string `json:"key"`
}

func GetTLSConfigFromAcmeJson(path string, domain string) (*tls.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var acme AcmeFile

	if err := json.Unmarshal(data, &acme); err != nil {
		return nil, err
	}

	for _, cert := range acme.LetsEncrypt.Certificates {
		if cert.Domain.Main != domain {
			continue
		}

		certPEM, err := base64.StdEncoding.DecodeString(cert.Certificate)
		if err != nil {
			return nil, err
		}

		keyPEM, err := base64.StdEncoding.DecodeString(cert.Key)
		if err != nil {
			return nil, err
		}

		tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
		if err != nil {
			return nil, err
		}

		return &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			MinVersion:   tls.VersionTLS12,
		}, nil
	}
	return nil, errors.New("No certificate found!")
}
