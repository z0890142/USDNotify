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
					var to string

					replyToken := event.ReplyToken
					name := strings.Split(message.Text, " ")[1]
					groupID := event.Source.GroupID

					userID := event.Source.UserID
					if groupID != "" {
						to = groupID
					} else {
						to = userID
					}
					err := DB.CheclUserExist(to)
					if err != nil {
						Log.Error(err)
						service.ReplyMessage(replyToken, "訂閱失敗", Log)
						return
					}
					err = DB.InsertSubscribeMember(name, to)
					if err != nil && strings.Contains(err.Error(), "sql: no rows in result set") {
						Log.Error(err)
						service.ReplyMessage(replyToken, "重複訂閱", Log)
						return
					} else if err != nil {
						Log.Error(err)
						service.ReplyMessage(replyToken, "訂閱失敗", Log)
						return
					}
					service.ReplyMessage(replyToken, "訂閱成功", Log)
				} else if strings.Contains(message.Text, "歷年") {
					replyToken := event.ReplyToken
					name := strings.Split(message.Text, " ")[1]
					queryType := strings.Split(message.Text, " ")[2]

					if queryType == "買價" {
						foreignCurrency.GetBuyInPriceRecord(name, replyToken, Log)
					} else if queryType == "賣價" {
						foreignCurrency.GetSellPriceRecord(name, replyToken, Log)
					} else {
						service.ReplyMessage(replyToken, "抱歉我聽不懂你在講什麼", Log)
					}

				} else {
					replyToken := event.ReplyToken
					foreignCurrency.GetNowPrice(message.Text, replyToken, Log)
				}

			}
		}

	}

}

func PictureTest(w http.ResponseWriter, r *http.Request) {
	Log, _ := Comman.LogInit("PictureTest", "USDNotify", logrus.DebugLevel)

	service.GetChart(1, "", Log)
}
