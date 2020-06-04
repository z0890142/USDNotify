package service

import (
	"USDNotify/helper/DB"
	"bytes"
	"os/exec"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

func GetChart(SN int, to string, Log *logrus.Entry) {
	displayName := DB.GetNameBySN(SN)
	//relative path of main.go or ABSOLUTE_PATH
	cmd := exec.Command("python3", "./python/main.py", displayName, strconv.Itoa(SN))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		Log.WithFields(logrus.Fields{
			"Error": stderr.String(),
		}).Error("GetChart Error")
		return
	} else {
		Log.Info(out.String())
	}

	bot.PushMessage(to, linebot.NewImageMessage("", ""))
}
