package main

import "os"

type Env struct {
	Domain       string
	SmtpDomain   string
	AcmeJsonPath string
}

func LoadEnv() *Env {
	env := &Env{}
	env.Domain = os.Getenv("DOMAIN")
	if env.Domain == "" {
		env.Domain = "localhost"
	}

	env.SmtpDomain = os.Getenv("SMTP_DOMAIN")
	if env.SmtpDomain == "" {
		env.SmtpDomain = "localhost"
	}

	env.AcmeJsonPath = os.Getenv("ACME_JSON_PATH")
	if env.AcmeJsonPath == "" {
		env.AcmeJsonPath = "/etc/letsencrypt/acme.json"
	}
	return env
}
