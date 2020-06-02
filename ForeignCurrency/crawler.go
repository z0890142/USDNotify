package foreignCurrency

import (
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/sirupsen/logrus"
)

func Crawler() {
	soup.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	}

	source, err := soup.Get("https://rate.bot.com.tw/xrt?Lang=zh-TW")
	if err != nil {

	}
	doc := soup.HTMLParse(source)
	for _, root := range doc.Find("table", "title", "牌告匯率").FindAll("tr") {

		tdArray := root.FindAll("td")
		if len(tdArray) == 0 {
			continue
		}
		name := tdArray[0].Find("div", "class", "print_show").Text()
		name = strings.Replace(name, " ", "", -1)
		name = strings.Replace(name, "\n", "", -1)

		for _, f := range ForeignCurrencyMap {
			if !strings.Contains(name, f.Name) {
				continue
			}
			buyInStr := tdArray[3].Text()
			sellStr := tdArray[4].Text()

			buyIn, _ := strconv.ParseFloat(buyInStr, 64)
			sell, _ := strconv.ParseFloat(sellStr, 64)
			Log.WithFields(logrus.Fields{
				"buyInStr": buyInStr,
				"sellStr":  sellStr,
				"buyIn":    buyIn,
				"sell":     sell,
			}).Info("Price info")
			f.Now_sell = sell
			f.Now_buyIn = buyIn
			f.UpdateTime = time.Now().Format("2006-01-02 15:04")
			f.LowestHandler(sell)
			f.HeigestHandler(buyIn)
			break
		}

	}
}
