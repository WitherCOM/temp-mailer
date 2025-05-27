package main

import (
	"log"

	"github.com/emersion/go-smtp"
	"github.com/gin-gonic/gin"
)

type MailerUri struct {
	Mailer string `uri:"mailer" binding:"required"`
}

func main() {
	smtpSignal := make(chan error)
	webSignal := make(chan error)
	storage := InitInMemoryStorage()
	env := LoadEnv()

	go func() {
		server := smtp.NewServer(&MailBackend{Domain: env.Domain, Storage: storage})
		server.Addr = ":25"
		server.AllowInsecureAuth = true
		log.Println("Starting server at", server.Addr)
		smtpSignal <- server.ListenAndServe()
	}()
	go func() {
		router := gin.Default()
		router.GET("/mail/:mailer", func(c *gin.Context) {
			var mailer MailerUri
			if err := c.ShouldBindUri(&mailer); err != nil {
				c.String(400, "Invalid mailer")
				return
			}
			if err := storage.AssignMail(mailer.Mailer); err != nil {
				c.String(500, "Failed to assign mail")
				return
			}
			c.String(200, mailer.Mailer+"@"+env.Domain)
		})
		router.GET("/mail/:mailer/content", func(c *gin.Context) {
			var mailer MailerUri
			if err := c.ShouldBindUri(&mailer); err != nil {
				c.String(400, "Invalid mailer")
				return
			}
			content, err := storage.GetMailContent(mailer.Mailer)
			if err != nil || content == "" {
				c.String(500, "Failed to get mail content")
				return
			}
			c.String(200, content)
		})
		webSignal <- router.Run(":8080")
	}()

	select {
	case err := <-smtpSignal:
		log.Fatal(err)
	case err := <-webSignal:
		log.Fatal(err)
	}
}
