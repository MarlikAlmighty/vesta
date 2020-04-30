package main

import (
	"github.com/MarlikAlmighty/vesta/bot"
	"github.com/MarlikAlmighty/vesta/models"
	"github.com/go-openapi/strfmt"
	"log"
	"os"
)

func main() {

	cfg := new(models.Config)

	webHook := os.Getenv("WEB_HOOK")
	cfg.WebHook = &webHook

	botToken := os.Getenv("BOT_TOKEN")
	cfg.BotToken = &botToken

	port := os.Getenv("PORT")
	cfg.Port = &port

	host := "0.0.0.0"
	cfg.Host = &host

	if err := cfg.Validate(strfmt.Default); err != nil {
		log.Fatalf("Error validation config: %v\n", err)
	}

	if err := bot.Run(cfg); err != nil {
		log.Fatalf("Error bot run: %v\n", err)
	}
}
