package main

import "os"

type Env struct {
	Domain       string
	AcmeJsonPath string
}

func LoadEnv() *Env {
	env := &Env{}
	env.Domain = os.Getenv("DOMAIN")
	if env.Domain == "" {
		env.Domain = "localhost"
	}
	env.AcmeJsonPath = os.Getenv("ACME_JSON_PATH")
	if env.AcmeJsonPath == "" {
		env.AcmeJsonPath = "localhost"
	}
	return env
}
