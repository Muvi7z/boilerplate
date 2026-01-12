package telegram

import (
	"bytes"
	"context"
	"embed"
	"github.com/Muvi7z/boilerplate/notification/internal/client/http"
	"github.com/Muvi7z/boilerplate/notification/internal/entity"
	"github.com/Muvi7z/boilerplate/platform/logger"
	"go.uber.org/zap"
	"text/template"
)

//go:embed templates/order_notification.tmpl
var templateFS embed.FS

type orderTemplateData struct {
	EventUuid       string
	OrderUuid       string
	UserUuid        string
	PaymentMethod   string
	TransactionUuid string
	BuildTimeSec    int64
}

var orderTemplate = template.Must(template.ParseFS(templateFS, "templates/order_notification.tmpl"))

type service struct {
	telegramClient http.TelegramClient
}

const chatID = 234586218

func (s *service) SendNotificationAssembly(ctx context.Context, event entity.ShipAssembledEvent) error {
	data := orderTemplateData{
		EventUuid:    event.EventUuid,
		OrderUuid:    event.OrderUuid,
		UserUuid:     event.UserUuid,
		BuildTimeSec: event.BuildTimeSec,
	}

	message, err := s.buildOrderMessage(data)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Telegram message sent to chat", zap.Int("chat_id", chatID), zap.String("message", message))
	return nil
}

func (s *service) SendNotificationPaid(ctx context.Context, event entity.OrderPaidEvent) error {

	data := orderTemplateData{
		EventUuid:       event.EventUuid,
		OrderUuid:       event.OrderUuid,
		UserUuid:        event.UserUuid,
		PaymentMethod:   event.PaymentMethod,
		TransactionUuid: event.TransactionUuid,
	}

	message, err := s.buildOrderMessage(data)
	if err != nil {
		return err
	}

	err = s.telegramClient.SendMessage(ctx, chatID, message)
	if err != nil {
		return err
	}

	logger.Info(ctx, "Telegram message sent to chat", zap.Int("chat_id", chatID), zap.String("message", message))
	return nil
}
func (s *service) buildOrderMessage(data orderTemplateData) (string, error) {
	var buf bytes.Buffer
	err := orderTemplate.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
