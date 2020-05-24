package DB

import (
	"fmt"
	"time"

	//前面加 _ 是為了只讓他執行init
	"USDNotify/model"

	_ "github.com/go-sql-driver/mysql" //前面加 _ 是為了只讓他執行init

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var db *sqlx.DB

func CreateDbConn(driveName string, dataSourceName string, Log *logrus.Entry) error {
	var err error
	// SetTls(Log)
	db, err = sqlx.Open(driveName, dataSourceName)
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("CreateDbConn error")
		return err
	}

	Log.Info("connnect success")

	return err
}

func GetForeignCurrencyList() (list []int, nameList []string, err error) {
	sqlString := "select SN,Name from ForeignCurrency"
	rows, err := db.Query(sqlString)
	if err != nil {
		err = fmt.Errorf("Get CurrentPrice : %v", err)
		return
	}
	for rows.Next() {
		var tmpSN int
		var tmpName string
		rows.Scan(&tmpSN, &tmpName)
		list = append(list, tmpSN)
		nameList = append(nameList, tmpName)
	}

	return
}

func GetForeignCurrencyRecord(SN int) (recordList model.ForeignCurrencyRecord, err error) {
	sqlString := "select 15_Lowest,15_Heigest,3Month_Lowest,3Month_Heigest,6Month_Lowest,6Month_Heigest," +
		"1Year_Lowest,1Year_Heigest,3Year_Lowest,3Year_Heigest,5Year_Lowest,5Year_Heigest from ForeignCurrencyRecord " +
		"where SN=?"
	db.QueryRow(sqlString, SN).Scan(
		&recordList.Lowest,
		&recordList.Heigest,
		&recordList.ThirdMonth_Lowest,
		&recordList.ThirdMonth_Heigest,
		&recordList.SixMonth_Lowest,
		&recordList.SixMonth_Heigest,
		&recordList.OneYear_Lowest,
		&recordList.OneYear_Heigest,
		&recordList.ThirdYear_Lowest,
		&recordList.ThirdYear_Heigest,
		&recordList.FiveYear_Lowest,
		&recordList.FiveYear_Heigest,
	)

	return
}

func SaveTodayPrice(name string, sellPrice float64, buyInPrice float64) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("SaveTodaySellPrice : %v", err)
	}
	date := time.Now().In(time.FixedZone("CST", 28800)).Format("2006-01-02")
	_, err = tx.Exec("insert into ForeignCurrencySellPrice(Name,Price,Date) values(?,?,?)",
		name, sellPrice, date)
	if err != nil {
		return fmt.Errorf("SaveTodaySellPrice : %v", err)
	}
	_, err = tx.Exec("insert into ForeignCurrencyBuyInPrice(Name,Price,Date) values(?,?,?)",
		name, buyInPrice, date)
	if err != nil {
		return fmt.Errorf("SaveTodayBuyInPrice : %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Commit : %v", err)
	}
	return nil
}

func GetSubscribeMember(name string) (userList []string, err error) {
	sqlString := "select UserId from Subscribe where name=?"
	rows, err := db.Query(sqlString, name)
	if err != nil {
		return userList, fmt.Errorf("GetSubscribeMember : %v", err)
	}
	for rows.Next() {
		var tmpUser string
		rows.Scan(&tmpUser)
		userList = append(userList, tmpUser)
	}

	return
}

func InsertSubscribeMember(name string, userID string) error {
	sqlString := "insert into Subscribe(Name,UserId) values(?,?)"
	_, err := db.Exec(sqlString, name, userID)
	if err != nil {
		return fmt.Errorf("InsertSubscribeMember : %v", err)
	}
	return nil

}
