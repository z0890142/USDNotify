package controller

import (
	foreignCurrency "USDNotify/ForeignCurrency"
	"USDNotify/helper/Comman"
	"USDNotify/helper/DB"
	service "USDNotify/services"
	"net/http"
	"strings"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

var bot *linebot.Client

func SetLineBot(_bot *linebot.Client) {
	bot = _bot
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	Log, _ := Comman.LogInit("CallbackHandler", "USDNotify", logrus.DebugLevel)
	events, err := bot.ParseRequest(r)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("lineBot.ParseRequest error")
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//member add group
				if strings.Contains(message.Text, "訂閱") {
					name := strings.Split(message.Text," ")[1]
					userID := event.Source.UserID
					err:=DB.CheclUserExist(userID)
					if err != nil {
						Log.Error(err)
						service.PushMessage(userID, "訂閱失敗", Log)
						return
					}
					err = DB.InsertSubscribeMember(name, userID)
					if err != nil {
						Log.Error(err)
						service.PushMessage(userID, "訂閱失敗", Log)
						return
					}
					service.PushMessage(userID, "訂閱成功", Log)
				} else {
					userID := event.Source.UserID
					foreignCurrency.GetNowPrice(message.Text, userID, Log)

				}

			}
		}

	}

}
