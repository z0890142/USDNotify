package main

import (
	foreignCurrency "USDNotify/ForeignCurrency"
	"USDNotify/controller"
	"USDNotify/helper/Comman"
	"USDNotify/helper/DB"
	service "USDNotify/services"

	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Log *logrus.Entry

func init() {
	Log, _ = Comman.LogInit("main", "USDNotify", logrus.DebugLevel)
	Log.Info("USDNotify version 0.0.0")
	InitConfig()
	DB.CreateDbConn("mysql", viper.GetString("DB.connectString"), Log)
	foreignCurrency.Init()
	Init_bot()
}

func Init_bot() {
	var err error
	channelSecret := viper.GetString("Line.ChannelSecret")
	channelAccessToken := viper.GetString("Line.ChannelAccessToken")
	bot, err := linebot.New(channelSecret, channelAccessToken)
	controller.SetLineBot(bot)
	service.SetLineBot(bot)

	if err != nil {
		Log.WithFields(logrus.Fields{
			"ChannelSecret":      viper.GetString("Line.ChannelSecret"),
			"ChannelAccessToken": viper.GetString("Line.ChannelAccessToken"),
		}).Error("new line bot error")
	}
}
func main() {

	port := "8088"
	addr := fmt.Sprintf(":%s", port)

	router := mux.NewRouter()
	router.HandleFunc("/callback", controller.CallbackHandler)
	router.HandleFunc("/PictureTest", controller.PictureTest).Methods("POST")
	// fs := http.FileServer(http.Dir("./static/.well-known/acme-challenge"))
	// router.PathPrefix("/.well-known/acme-challenge/").Handler(http.StripPrefix("/.well-known/acme-challenge/", fs))

	fs := http.FileServer(http.Dir("./static/picture"))
	router.PathPrefix("/picture/").Handler(http.StripPrefix("/picture/", fs))

	// err := http.ListenAndServe(addr, router)

	err := http.ListenAndServeTLS(addr, "./static/ssl/bundle.crt", "./static/ssl/private.key", router)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()         // read in environment variables that match
	viper.SetEnvPrefix("gorush") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err == nil {
		Log.Info("Using config file:", viper.ConfigFileUsed())
	}

}
