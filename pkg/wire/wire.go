//go:build wireinject
// +build wireinject

package wire_excel

import (
	"github.com/google/wire"
	"github.com/ryvasa/go-super-farmer-mail-service/cmd/app"
	mail_handler "github.com/ryvasa/go-super-farmer-mail-service/internal/handler"
	mail_usecase "github.com/ryvasa/go-super-farmer-mail-service/internal/usecase"
	"github.com/ryvasa/go-super-farmer-mail-service/pkg/env"
	"github.com/ryvasa/go-super-farmer-mail-service/pkg/messages"
)

var allSet = wire.NewSet(
	env.LoadEnv,
	messages.NewRabbitMQ,
	mail_usecase.NewMailUsecase,
	mail_handler.NewMailHandler,
	app.NewApp,
)

func InitializeMailApp() (*app.MailApp, error) {
	wire.Build(allSet)
	return nil, nil
}
