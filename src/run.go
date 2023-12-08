package main

import (
	"github.com/msw-x/moon/app"
	"github.com/msw-x/moon/telegram"
	"github.com/msw-x/moon/ulog"
)

func run(conf Conf) {
	conf.FromEnv()

	bot := telegram.NewAlertBot(conf.TelegramBotToken, conf.TelegramBotChatId, version)
	bot.Startup()
	defer bot.Shutdown(0)

	ulog.SetHook(func(m ulog.Message) {
		//tbot.SendLog(m)
	})

	/*
		db := db.New(db.Options{
			User:          conf.DbUser,
			Pass:          conf.DbPass,
			Host:          conf.DbHost,
			Name:          conf.DbName,
			Timeout:       time.Second * 8,
			MaxConnFactor: 2,
			MinOpenConns:  4,
			Strict:        true,
			Insecure:      true,
			LogErrors:     conf.DbLogErrors,
			LogQueries:    conf.DbLogQueries,
		})
		defer db.Close()
		db.Await(time.Second * 4)
	*/

	app.WaitInterrupt()
}
