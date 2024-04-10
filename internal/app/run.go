package app

import (
	"module/internal/database"
	"module/internal/server"

	log "github.com/sirupsen/logrus"
)

func RunApp() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// подключение к пулу соединений
	var temp database.PostgressCon
	temp.NewConnection()
	database.GlobalInstance = &temp

	// переподключение к бд
	go temp.PingReconnect()

	// миграция
	// go database.MigrateStart()

	// запуск сервера. внутри указан порт 8888
	server.MainServer()

}
