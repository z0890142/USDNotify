package service

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

var bot *linebot.Client

func SetLineBot(_bot *linebot.Client) {
	bot = _bot
}

func PushMessage(to string, payload string, Log *logrus.Entry) {
	if _, err := bot.PushMessage(to, linebot.NewTextMessage(payload)).Do(); err != nil {
		Log.Error(err)
	}
}

func ReplyMessage(replyToken string, payload string, Log *logrus.Entry) {
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(payload)).Do(); err != nil {
		Log.Error(err)
	}
}
