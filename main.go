package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token: os.Getenv("TELE_TOKEN"),
		Poller: &tele.LongPoller{
			Timeout: 10 * time.Second,
		},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! this is a simple bot that echoes your messages. Send me any text and I will repeat it back to you.")
	})

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello! How are you doing today?")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send(c.Text())
	})

	b.Start()
}
