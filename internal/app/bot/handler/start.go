package handler

import (
	"github.com/abg216te/tgbt/internal/infrastructure/telegram"
	"github.com/abg216te/tgbt/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
	bot     *telegram.Bot
	service *service.UserService
}

func NewStartHandler(bot *telegram.Bot, s *service.UserService) *StartHandler {
	return &StartHandler{bot: bot, service: s}
}

func (h *StartHandler) Handle(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Text == "/start" {
		user := update.Message.From

		_ = h.service.Register(user.ID, user.UserName)

		h.bot.Send(update.Message.Chat.ID, "Welcome!")
	}
}
