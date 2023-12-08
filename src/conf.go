package main

import (
	"os"
)

type Conf struct {
	LogTimezone       string
	UiHost            string
	DbHost            string
	DbUser            string
	DbPass            string
	DbName            string
	DbLogErrors       bool
	DbLogQueries      bool
	DbLock            bool
	DbRollback        bool
	DbPreviewDown     bool
	UiAddr            string
	CertDir           string
	CertHost          string
	TelegramBotToken  string
	TelegramBotChatId string
}

func (o *Conf) FromEnv() {
	fromEnv("UI_HOST", &o.UiHost)
	fromEnv("DB_HOST", &o.DbHost)
	fromEnv("DB_NAME", &o.DbName)
	fromEnv("DB_USER", &o.DbUser)
	fromEnv("DB_PASS", &o.DbPass)
	fromEnv("CERT_HOST", &o.CertHost)
	fromEnv("TELEGRAM_TOKEN", &o.TelegramBotToken)
	fromEnv("TELEGRAM_CHAT_ID", &o.TelegramBotChatId)
}

func fromEnv(name string, param *string) {
	if v, ok := os.LookupEnv(name); ok {
		*param = v
	}
}
