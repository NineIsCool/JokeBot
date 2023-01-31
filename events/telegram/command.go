package telegram

import (
	"log"
	"strings"
)

const (
	RndJokeCmd = "/rndjoke"
	HelpCmd    = "/help"
	StartCmd   = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	case RndJokeCmd:
		return p.sendJoke(chatID)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func (p *Processor) sendJoke(chatID int) error {
	return p.tg.SendMessage(chatID, p.storage.RndJoke())
}
