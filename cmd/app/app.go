package app

import (
	mail_handler "github.com/ryvasa/go-super-farmer-mail-service/internal/handler"
	"github.com/ryvasa/go-super-farmer-mail-service/pkg/env"
	"github.com/ryvasa/go-super-farmer-mail-service/pkg/logrus"
	"github.com/ryvasa/go-super-farmer-mail-service/pkg/messages"
)

type MailApp struct {
	Env      *env.Env
	RabbitMQ messages.RabbitMQ
	Handler  mail_handler.MailHandler
}

func NewApp(
	env *env.Env,
	rabbitMQ messages.RabbitMQ,
	handler mail_handler.MailHandler,
) *MailApp {
	err := handler.ConsumerHandler()
	if err != nil {
		logrus.Log.Fatalf("failed to initiate consumer handler: %v", err)
	}
	return &MailApp{
		Env:      env,
		RabbitMQ: rabbitMQ,
		// Handler:  handler,
	}
}
