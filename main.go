package main

import (
	tgClient "JokeBot/clients/telegram"
	"JokeBot/consumer/event-consumer"
	"JokeBot/events/telegram"
	"JokeBot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	log.Println("service started")
	consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

func mustToken() string {
	toke := flag.String("tg-bot-token", "", "token for access to telegram bot")
	flag.Parse()
	if *toke == "" {
		log.Fatal("token is not specified")
	}
	return *toke
}
