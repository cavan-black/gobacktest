package main

import (
	"fmt"

	"github.com/cavan-black/gobacktest"
	"github.com/cavan-black/gobacktest/data"
	"github.com/cavan-black/gobacktest/strategy"
)

func main() {
	// initiate a new backtester
	test := gobacktest.New()
	portfolio := gobacktest.Portfolio{}
	portfolio.SetInitialCash(10000)

	sizeManager := &gobacktest.Size{DefaultSize: 100, DefaultValue: 1000}
	portfolio.SetSizeManager(sizeManager)

	riskManager := &gobacktest.Risk{}
	portfolio.SetRiskManager(riskManager)

	test.SetPortfolio(&portfolio)
	// define and load symbols
	symbols := []string{"BAS.DE"}
	test.SetSymbols(symbols)

	// create a data provider and load the data into the backtest
	data := &data.BarEventFromCSVFile{FileDir: "../testdata/bar/"}
	data.Load(symbols)
	test.SetData(data)
	fmt.Print(data)
	// choose a strategy
	strategy := strategy.BuyAndHold()

	// create an asset and append it to the strategy
	strategy.SetChildren(gobacktest.NewAsset("BAS.DE"))

	// load the strategy into the backtest
	test.SetStrategy(strategy)

	// run the backtest
	test.Run()

	// print the results of the test
	test.Stats().PrintResult()
}
