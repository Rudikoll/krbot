package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {
	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		log.Fatal("TELE_TOKEN not set")
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	kbot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	kbot.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send("Привіт! Ти написав: " + c.Text())
	})

	kbot.Start()
}
