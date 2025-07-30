package main

import "os"

type Env struct {
	Domain string
}

func LoadEnv() *Env {
	env := &Env{}
	env.Domain = os.Getenv("DOMAIN")
	if env.Domain == "" {
		env.Domain = "localhost"
	}
	return env
}
