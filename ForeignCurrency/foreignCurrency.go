package foreignCurrency

import (
	"USDNotify/helper/Comman"
	"USDNotify/helper/DB"
	service "USDNotify/services"
	"fmt"
	"strings"

	"github.com/robfig/cron"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry
var ForeignCurrencyMap map[int]*ForeignCurrency
var masterCron *cron.Cron

type ForeignCurrency struct {
	SN               int
	Name             string
	Subscribe_Number int
	ForeignCurrencyRecord
	Log           *logrus.Entry
	Today_Lowest  float64
	Today_Heigest float64
	Now_sell      float64
	Now_buyIn     float64
	UpdateTime    string
}

type ForeignCurrencyRecord struct {
	Lowest             float64
	Heigest            float64
	ThirdMonth_Lowest  float64
	ThirdMonth_Heigest float64
	SixMonth_Lowest    float64
	SixMonth_Heigest   float64
	OneYear_Lowest     float64
	OneYear_Heigest    float64
	ThirdYear_Lowest   float64
	ThirdYear_Heigest  float64
	FiveYear_Lowest    float64
	FiveYear_Heigest   float64
}

func init() {
	Log, _ = Comman.LogInit("service", "USDNotify", logrus.DebugLevel)
	ForeignCurrencyMap = make(map[int]*ForeignCurrency)
	masterCron = cron.New()
	masterCron.AddFunc("0 0/5 9-17 * * *", Crawler)
	masterCron.Start()

}

func (f *ForeignCurrency) SaveTodayPrice() {
	err := DB.SaveTodayPrice(f.SN, f.Today_Lowest, f.Today_Heigest)
	if err != nil {
		f.Log.WithFields(logrus.Fields{
			"Error": err,
		}).Error("SaveTodayPrice Error")
		return
	}
}

func (f *ForeignCurrency) LowestHandler(sell float64) {
	var msg string
	if f.FiveYear_Lowest > sell {
		f.FiveYear_Lowest = sell
		f.ThirdYear_Lowest = sell
		f.OneYear_Lowest = sell
		f.SixMonth_Lowest = sell
		f.ThirdMonth_Lowest = sell
		f.Lowest = sell
		f.Today_Lowest = sell
		msg = f.Name + "銀行賣價已達五年內最低價"
	} else if f.ThirdMonth_Lowest > sell {
		f.ThirdYear_Lowest = sell
		f.OneYear_Lowest = sell
		f.SixMonth_Lowest = sell
		f.ThirdMonth_Lowest = sell
		f.Lowest = sell
		f.Today_Lowest = sell
		msg = f.Name + "銀行賣價已達三年內最低價"
	} else if f.OneYear_Lowest > sell {
		f.OneYear_Lowest = sell
		f.SixMonth_Lowest = sell
		f.ThirdMonth_Lowest = sell
		f.Lowest = sell
		f.Today_Lowest = sell
		msg = f.Name + "銀行賣價已達一年內最低價"
	} else if f.SixMonth_Lowest > sell {
		f.SixMonth_Lowest = sell
		f.ThirdMonth_Lowest = sell
		f.Lowest = sell
		f.Today_Lowest = sell
		msg = f.Name + "銀行賣價已達半年內最低價"
	} else if f.ThirdMonth_Lowest > sell {
		f.ThirdMonth_Lowest = sell
		f.Lowest = sell
		f.Today_Lowest = sell
		msg = f.Name + "銀行賣價已達三個月內最低價"
	} else if f.Lowest > sell {
		f.Lowest = sell
		f.Today_Lowest = sell
		msg = f.Name + "銀行賣價已達半個月內最低價"
	} else if f.Today_Lowest > sell || f.Today_Lowest == 0 {
		f.Today_Lowest = sell
	}

	if msg == "" {
		return
	}

	userList, err := DB.GetSubscribeMember(f.SN)
	if err != nil {
		f.Log.WithFields(logrus.Fields{
			"Error": err,
		}).Error("GetSubscribeMember Error")
		return
	}

	for _, user := range userList {
		service.PushMessage(user, msg, f.Log)
	}
}

func (f *ForeignCurrency) HeigestHandler(buyin float64) {
	var msg string
	if f.FiveYear_Heigest < buyin {
		f.FiveYear_Heigest = buyin
		f.ThirdYear_Heigest = buyin
		f.OneYear_Heigest = buyin
		f.SixMonth_Heigest = buyin
		f.ThirdMonth_Heigest = buyin
		f.Heigest = buyin
		f.Today_Heigest = buyin
		msg = f.Name + "銀行買價已達五年內最高價"
	} else if f.ThirdMonth_Heigest < buyin {
		f.ThirdYear_Heigest = buyin
		f.OneYear_Heigest = buyin
		f.SixMonth_Heigest = buyin
		f.ThirdMonth_Heigest = buyin
		f.Heigest = buyin
		f.Today_Heigest = buyin
		msg = f.Name + "銀行買價已達三年內最高價"
	} else if f.OneYear_Heigest < buyin {
		f.OneYear_Heigest = buyin
		f.SixMonth_Heigest = buyin
		f.ThirdMonth_Heigest = buyin
		f.Heigest = buyin
		f.Today_Heigest = buyin
		msg = f.Name + "銀行買價已達一年內最高價"
	} else if f.SixMonth_Heigest < buyin {
		f.SixMonth_Heigest = buyin
		f.ThirdMonth_Heigest = buyin
		f.Heigest = buyin
		f.Today_Heigest = buyin
		msg = f.Name + "銀行買價已達半年內最高價"
	} else if f.ThirdMonth_Heigest < buyin {
		f.ThirdMonth_Heigest = buyin
		f.Heigest = buyin
		f.Today_Heigest = buyin
		msg = f.Name + "銀行買價已達三個月內最高價"
	} else if f.Heigest < buyin {
		f.Heigest = buyin
		f.Today_Heigest = buyin
		msg = f.Name + "銀行買價已達半個月內最高價"
	} else if f.Today_Heigest < buyin {
		f.Today_Heigest = buyin
	}

	if msg == "" {
		return
	}

	userList, err := DB.GetSubscribeMember(f.SN)
	if err != nil {
		f.Log.WithFields(logrus.Fields{
			"Error": err,
		}).Error("GetSubscribeMember Error")
		return
	}

	for _, user := range userList {
		service.PushMessage(user, msg, f.Log)
	}
}
func Init() {
	currencyList, nameList, err := DB.GetForeignCurrencyList()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetForeignCurrencyList Error")
	}
	for index, currency := range currencyList {
		record, err := DB.GetForeignCurrencyRecord(currency)
		if err != nil {
			Log.WithFields(logrus.Fields{
				"error": err,
			}).Error("GetForeignCurrencyRecord Error")
		}
		_Log, _ := Comman.LogInit(nameList[index], "USDNotify", logrus.DebugLevel)

		tmpForeignCurrency := &ForeignCurrency{
			Name:             nameList[index],
			SN:               currency,
			Subscribe_Number: 0,
			Log:              _Log,
		}

		tmpForeignCurrency.Heigest = record.Heigest
		tmpForeignCurrency.Lowest = record.Lowest
		tmpForeignCurrency.ThirdYear_Heigest = record.ThirdYear_Heigest
		tmpForeignCurrency.ThirdYear_Lowest = record.ThirdYear_Lowest
		tmpForeignCurrency.SixMonth_Heigest = record.SixMonth_Heigest
		tmpForeignCurrency.SixMonth_Lowest = record.SixMonth_Lowest
		tmpForeignCurrency.OneYear_Heigest = record.OneYear_Heigest
		tmpForeignCurrency.OneYear_Lowest = record.OneYear_Lowest
		tmpForeignCurrency.ThirdMonth_Heigest = record.ThirdMonth_Heigest
		tmpForeignCurrency.ThirdMonth_Lowest = record.ThirdMonth_Lowest
		tmpForeignCurrency.FiveYear_Heigest = record.FiveYear_Heigest
		tmpForeignCurrency.FiveYear_Lowest = record.FiveYear_Lowest

		masterCron.AddFunc("0 0 16 * * *", tmpForeignCurrency.SaveTodayPrice)
		ForeignCurrencyMap[tmpForeignCurrency.SN] = tmpForeignCurrency
	}
}

func GetNowPrice(name string, to string, Log *logrus.Entry) {
	for _, foreignCurrency := range ForeignCurrencyMap {
		if strings.Contains(foreignCurrency.Name, strings.ToUpper(name)) {
			msg := foreignCurrency.Name + "\n 銀行即期賣價 : " + fmt.Sprintf("%v", foreignCurrency.Now_sell) +
				"\n 銀行即期買價 : " + fmt.Sprintf("%v", foreignCurrency.Now_buyIn) +
				"\n 更新時間 : " + foreignCurrency.UpdateTime
			service.PushMessage(to, msg, Log)
		}
	}
}
