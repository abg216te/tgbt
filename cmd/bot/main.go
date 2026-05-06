package main

import (
	"fmt"
	"log"

	"github.com/abg216te/tgbt/internal/app/bot/handler"
	"github.com/abg216te/tgbt/internal/config"
	"github.com/abg216te/tgbt/internal/infrastructure/telegram"
	"github.com/abg216te/tgbt/internal/repository/memory"
	"github.com/abg216te/tgbt/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot
func main() {
	сfg, err := config.Load()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	token := сfg.Token

	fmt.Println("Token:", token)

	repo := memory.NewUserRepo()
	service := service.NewUserService(repo)

	bot, err := telegram.NewBot(token)
	if err != nil {
		log.Fatal("Error creating bot:", err)
		return
	}

	startHandler := handler.NewStartHandler(bot, service)

	bot.Start(func(update tgbotapi.Update) {
		startHandler.Handle(update)
	})
}
