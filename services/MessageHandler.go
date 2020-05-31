package service

import (
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

var bot *linebot.Client
var foreignCurrencyMap *map[int]*cron.Cron

func SetLineBot(_bot *linebot.Client) {
	bot = _bot
}

func PushMessage(to string, payload string, Log *logrus.Entry) {
	if _, err := bot.PushMessage(to, linebot.NewTextMessage(payload)).Do(); err != nil {
		Log.Error(err)
	}
}

func SetForeignCurrencyMap(_foreignCurrencyMap *map[int]*cron.Cron) {
	foreignCurrencyMap = _foreignCurrencyMap
}

func GetNowPrice(name string) {
	for _, foreignCurrency := range *foreignCurrencyMap {
		if strings.Contains(foreignCurrency.Name, name) {
			msg := foreignCurrency.Name + "\n 銀行即期賣價 : " + fmt.Sprintf("%v", foreignCurrency.now_sell) + "\n 銀行即期買價 : " + fmt.Sprintf("%v", foreignCurrency.now_buyin)

			if _, err := bot.PushMessage(to, linebot.NewTextMessage(msg)).Do(); err != nil {
				Log.Error(err)
			}
		}
	}
}
