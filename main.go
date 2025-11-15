package main

import (
	tgClient "bot/clients/telegram"
	eventconsumer "bot/consumer/event-consumer"
	"bot/events/telegram"
	"bot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.bot"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("serive is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is specified")
	}

	return *token
}
