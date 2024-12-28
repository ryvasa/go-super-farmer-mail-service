package main

import (
	_ "github.com/lib/pq"
	"github.com/ryvasa/go-super-farmer-mail-service/pkg/logrus"
	wire_excel "github.com/ryvasa/go-super-farmer-mail-service/pkg/wire"
)

func main() {
	app, err := wire_excel.InitializeMailApp()
	defer app.RabbitMQ.Close()
	if err != nil {
		logrus.Log.Fatalf("failed to initialize app: %v", err)
	}
}
