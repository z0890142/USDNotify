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
	DisplayName      string
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

	Lowest_Date             string
	Heigest_Date            string
	ThirdMonth_Lowest_Date  string
	ThirdMonth_Heigest_Date string
	SixMonth_Lowest_Date    string
	SixMonth_Heigest_Date   string
	OneYear_Lowest_Date     string
	OneYear_Heigest_Date    string
	ThirdYear_Lowest_Date   string
	ThirdYear_Heigest_Date  string
	FiveYear_Lowest_Date    string
	FiveYear_Heigest_Date   string
}

func init() {
	Log, _ = Comman.LogInit("service", "USDNotify", logrus.DebugLevel)
	ForeignCurrencyMap = make(map[int]*ForeignCurrency)
	masterCron = cron.New()
	masterCron.AddFunc("0 0/10 9-16 * * *", Crawler)
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
	f.Today_Lowest = -99
	f.Today_Heigest = -99

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
	} else if f.ThirdYear_Lowest > sell {
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
	}
	if f.Today_Lowest > sell || f.Today_Lowest == -99 {
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
	} else if f.ThirdYear_Heigest < buyin {
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
	}

	if f.Today_Heigest < buyin || f.Today_Heigest == -99 {
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
	currencyList, nameList, dispalyNameList, err := DB.GetForeignCurrencyList()
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
			DisplayName:      dispalyNameList[index],
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

		tmpForeignCurrency.Lowest_Date = record.Lowest_Date
		tmpForeignCurrency.Heigest_Date = record.Heigest_Date
		tmpForeignCurrency.ThirdMonth_Heigest_Date = record.ThirdMonth_Heigest_Date
		tmpForeignCurrency.ThirdMonth_Lowest_Date = record.ThirdMonth_Lowest_Date
		tmpForeignCurrency.SixMonth_Lowest_Date = record.SixMonth_Lowest_Date
		tmpForeignCurrency.SixMonth_Heigest_Date = record.SixMonth_Heigest_Date
		tmpForeignCurrency.OneYear_Heigest_Date = record.OneYear_Heigest_Date
		tmpForeignCurrency.OneYear_Lowest_Date = record.OneYear_Lowest_Date
		tmpForeignCurrency.ThirdYear_Heigest_Date = record.ThirdYear_Heigest_Date
		tmpForeignCurrency.ThirdYear_Lowest_Date = record.ThirdYear_Lowest_Date
		tmpForeignCurrency.FiveYear_Heigest_Date = record.FiveYear_Heigest_Date
		tmpForeignCurrency.FiveYear_Lowest_Date = record.FiveYear_Lowest_Date

		masterCron.AddFunc("0 0 17 * * *", tmpForeignCurrency.SaveTodayPrice)
		ForeignCurrencyMap[tmpForeignCurrency.SN] = tmpForeignCurrency
	}
}

func GetNowPrice(name string, to string, Log *logrus.Entry) {
	for _, foreignCurrency := range ForeignCurrencyMap {
		if strings.Contains(foreignCurrency.DisplayName, strings.ToUpper(name)) {
			msg := foreignCurrency.DisplayName + "\n 銀行即期賣價 : " + fmt.Sprintf("%v", foreignCurrency.Now_sell) +
				"\n 銀行即期買價 : " + fmt.Sprintf("%v", foreignCurrency.Now_buyIn) +
				"\n 更新時間 : " + foreignCurrency.UpdateTime
			service.ReplyMessage(to, msg, Log)
			service.GetChart(foreignCurrency.SN, to, Log)

		}
	}
}

func GetBuyInPriceRecord(name string, to string, Log *logrus.Entry) {
	for _, foreignCurrency := range ForeignCurrencyMap {
		if strings.Contains(foreignCurrency.DisplayName, strings.ToUpper(name)) {
			msg := "銀行五年最高買價 : " + fmt.Sprintf("%v", foreignCurrency.FiveYear_Heigest) +
				" 日期 : " + foreignCurrency.FiveYear_Heigest_Date +
				"\n 銀行三年最高買價 : " + fmt.Sprintf("%v", foreignCurrency.ThirdYear_Heigest) +
				" 日期 : " + foreignCurrency.ThirdYear_Heigest_Date +
				"\n 銀行一年最高買價 : " + fmt.Sprintf("%v", foreignCurrency.OneYear_Heigest) +
				" 日期 : " + foreignCurrency.OneYear_Heigest_Date +
				"\n 銀行半年最高買價 : " + fmt.Sprintf("%v", foreignCurrency.SixMonth_Heigest) +
				" 日期 : " + foreignCurrency.SixMonth_Heigest_Date +
				"\n 銀行三個月最高買價 : " + fmt.Sprintf("%v", foreignCurrency.ThirdMonth_Heigest) +
				" 日期 : " + foreignCurrency.ThirdMonth_Heigest_Date +
				"\n 銀行15日最高買價 : " + fmt.Sprintf("%v", foreignCurrency.Heigest) +
				" 日期 : " + foreignCurrency.Heigest_Date

			service.ReplyMessage(to, msg, Log)

		}
	}
}

func GetSellPriceRecord(name string, to string, Log *logrus.Entry) {
	for _, foreignCurrency := range ForeignCurrencyMap {
		if strings.Contains(foreignCurrency.DisplayName, strings.ToUpper(name)) {
			msg := "銀行五年最低買價 : " + fmt.Sprintf("%v", foreignCurrency.FiveYear_Lowest) +
				" 日期 : " + foreignCurrency.FiveYear_Lowest_Date +
				"\n 銀行三年最低買價 : " + fmt.Sprintf("%v", foreignCurrency.ThirdYear_Lowest) +
				" 日期 : " + foreignCurrency.ThirdYear_Lowest_Date +
				"\n 銀行一年最低買價 : " + fmt.Sprintf("%v", foreignCurrency.OneYear_Lowest) +
				" 日期 : " + foreignCurrency.OneYear_Lowest_Date +
				"\n 銀行半年最低買價 : " + fmt.Sprintf("%v", foreignCurrency.SixMonth_Lowest) +
				" 日期 : " + foreignCurrency.SixMonth_Lowest_Date +
				"\n 銀行三個月最低買價 : " + fmt.Sprintf("%v", foreignCurrency.ThirdMonth_Lowest) +
				" 日期 : " + foreignCurrency.ThirdMonth_Lowest_Date +
				"\n 銀行15日最低買價 : " + fmt.Sprintf("%v", foreignCurrency.Lowest) +
				" 日期 : " + foreignCurrency.Lowest_Date

			service.ReplyMessage(to, msg, Log)

		}
	}
}
