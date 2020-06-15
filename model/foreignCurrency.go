package model

import "github.com/sirupsen/logrus"

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
