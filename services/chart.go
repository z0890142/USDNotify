package service

import (
	"USDNotify/helper/DB"
	"os/exec"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

func GetChart(SN int, to string, Log *logrus.Entry) {
	displayName := DB.GetNameBySN(SN)
	cmd := exec.Command("python", "./main.py", displayName, strconv.Itoa(SN))
	_, err := cmd.Output()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"Error": err,
		}).Error("GetChart Error")
		return
	}

	bot.PushMessage(to, linebot.NewImageMessage("", ""))
}
