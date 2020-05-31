package service

import(
	"os"
	"github.com/wcharczuk/go-chart"
	"USDNotify/helper/DB"
	"fmt"

)


func Get3MonthChart(SN int){
	yValues,xValues,err:=DB.Get3MonthSellPrice(SN)
	if err!=nil{
		fmt.Println(err.Error())
	}
	priceSeries := chart.TimeSeries{
		Name: "SPY",
		Style: chart.Style{
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xValues,
		YValues: yValues,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			TickPosition: chart.TickPositionBetweenTicks,
		},
		YAxis: chart.YAxis{
			Range: &chart.ContinuousRange{
				Max: 40.0,
				Min: 0,
			},
		},
		Series: []chart.Series{
			priceSeries,
		},
	}
	
	f, _ := os.Create("./static/picture/output.jpg")
	defer f.Close()
	graph.Render(chart.PNG, f)
}