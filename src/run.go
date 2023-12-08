package main

import (
	"embed"
	"time"

	"github.com/msw-x/moon/app"
	"github.com/msw-x/moon/db"
	"github.com/msw-x/moon/telegram"
	"github.com/msw-x/moon/ulog"

	"biton/core"
	"biton/webapi"
)

//go:embed migrations/*.sql
var sqlMigrations embed.FS

func run(conf Conf) {
	conf.FromEnv()

	bot := telegram.NewAlertBot(conf.TelegramBotToken, conf.TelegramBotChatId, version)
	bot.Startup()
	defer bot.Shutdown(0)

	ulog.SetHook(func(m ulog.Message) {
		bot.SendLog(m)
	})

	app.WaitInterrupt()
	return

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

	if db.Migrator().Exec(sqlMigrations, conf.DbLock, conf.DbRollback, conf.DbPreviewDown) {
		core := core.New(db)
		defer core.Close()
		core.Run()
		api := webapi.New(webapi.Options{
			UiAddr:   conf.UiAddr,
			CertDir:  conf.CertDir,
			CertHost: conf.CertHost,
			LogHttp:  true,
		}, core, version)
		defer api.Close()
	}

	app.WaitInterrupt()
}
