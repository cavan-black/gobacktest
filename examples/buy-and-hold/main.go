package main

import (
	"fmt"

	gbt "github.com/cavan-black/gobacktest"
	"github.com/cavan-black/gobacktest/algo"
	"github.com/cavan-black/gobacktest/data"
)

func main() {
	// initiate new backtester
	test := gbt.New()

	// define and load symbols
	var symbols = []string{"SDF.DE"}
	test.SetSymbols(symbols)

	// create data provider and load data into the backtest
	data := &data.BarEventFromCSVFile{FileDir: "../testdata/bar/"}
	data.Load(symbols)
	test.SetData(data)

	// create a new strategy with an algo stack and load into the backtest
	strategy := gbt.NewStrategy("basic")
	strategy.SetAlgo(
		algo.RunDaily(),          // run on beginning of each year
		algo.CreateSignal("buy"), // always create a buy signal on a data event
	)

	// create an asset and append to strategy
	strategy.SetChildren(gbt.NewAsset("SDF.DE"))

	// load the strategy into the backtest
	test.SetStrategy(strategy)

	// run the backtest
	err := test.Run()
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	// print the result of the test
	test.Stats().PrintResult()
}
