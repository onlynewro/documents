// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jroimartin/gocui"
)

type Upbit_Ticker struct {
	Market                string  `json:"market"`
	Trade_date            string  `json:"trade_date"`
	Trade_time            string  `json:"trade_time"`
	Trade_date_kst        string  `json:"trade_date_kst"`
	Trade_time_kst        string  `json:"trade_time_kst"`
	Opening_price         float32 `json:"onening_price"`
	High_price            float32 `json:"high_price"`
	Low_price             float32 `json:"low_price"`
	Trade_price           float32 `json:"trade_price"`
	Prev_closing_price    float32 `json:"prev_closing_price"`
	Change                string  `json:"change"`
	Change_price          float32 `json:"change_price"`
	Change_rate           float32 `json:"change_rate"`
	Signed_change_price   float32 `json:"signed_change_price"`
	Signed_change_rate    float32 `json:"signed_change_rate"`
	Trade_volume          float32 `json:"trade_volume"`
	Acc_trade_price       float32 `json:"acc_trade_price"`
	Acc_trade_price_24h   float32 `json:"acc_trade_price_24h"`
	Acc_trade_volume      float32 `json:"acc_trade_volume"`
	Acc_trade_volume_24h  float32 `json:"acc_trade_volume_24h"`
	Highest_52_week_price float32 `json:"highest_52_week_price"`
	Highest_52_week_date  string  `json:"highest_52_week_date"`
	Lowest_52_week_price  float32 `json:"lowest_52_week_price"`
	Lowest_52_week_date   string  `json:"lowest_52_week_date"`
	Timestamp             int64   `json:"timestamp"`
}

type Posts []Upbit_Ticker

type Orderbook_unit struct {
	Ask_price float32 `json:"ask_price"`
	Bid_price float32 `json:"bid_price"`
	Ask_size  float32 `json:"ask_size"`
	Bid_size  float32 `json:"bid_size"`
}

var Orderbook_units []Orderbook_unit

type Upbit_Orderbook struct {
	Market          string           `json:"KRW-BTC"`
	Timestamp       int64            `json:"timestamp"`
	Total_ask_size  float32          `json:"total_ask_size"`
	Total_bid_size  float32          `json:"total_bid_size"`
	Orderbook_units []Orderbook_unit `json:"orderbook_units"`
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 1, 1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true

		//		var up Upbit_Ticker

		var up Posts

		resp, err := http.Get("https://api.upbit.com/v1/ticker?markets=KRW-BTC")
		if err != nil {
			log.Panicln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(body, &up)
		if err != nil {
			fmt.Fprintln(v, "error:", err)
		}

		fmt.Fprintf(v, "Market:%s\n", up[0].Market)
		fmt.Fprintf(v, "trade_date:%s\n", up[0].Trade_date)
		fmt.Fprintf(v, "trade_time:%s\n", up[0].Trade_time)
		fmt.Fprintf(v, "trade_date_kst:%s\n", up[0].Trade_date_kst)
		fmt.Fprintf(v, "trade_time_kst:%s\n", up[0].Trade_time_kst)
		fmt.Fprintf(v, "onening_price:%f\n", up[0].Opening_price)
		fmt.Fprintf(v, "high_price:%f\n", up[0].High_price)
		fmt.Fprintf(v, "low_price:%f\n", up[0].Low_price)
		fmt.Fprintf(v, "trade_price:%f\n", up[0].Trade_price)
		fmt.Fprintf(v, "prev_closing_price:%f\n", up[0].Prev_closing_price)
		fmt.Fprintf(v, "change:%s\n", up[0].Change)
		fmt.Fprintf(v, "change_price:%f\n", up[0].Change_price)
		fmt.Fprintf(v, "change_rate:%f\n", up[0].Change_rate)
		fmt.Fprintf(v, "signed_change_price:%f\n", up[0].Signed_change_price)
		fmt.Fprintf(v, "signed_change_rate:%f\n", up[0].Signed_change_rate)
		fmt.Fprintf(v, "trade_volume:%f\n", up[0].Trade_volume)
		fmt.Fprintf(v, "acc_trade_price:%f\n", up[0].Acc_trade_price)
		fmt.Fprintf(v, "acc_trade_price_24h:%f\n", up[0].Acc_trade_price_24h)
		fmt.Fprintf(v, "acc_trade_volume:%f\n", up[0].Acc_trade_volume)
		fmt.Fprintf(v, "acc_trade_volume_24h:%f\n", up[0].Acc_trade_volume_24h)
		fmt.Fprintf(v, "highest_52_week_price:%f\n", up[0].Highest_52_week_price)
		fmt.Fprintf(v, "highest_52_week_date:%s\n", up[0].Highest_52_week_date)
		fmt.Fprintf(v, "lowest_52_week_price%f\n", up[0].Lowest_52_week_price)
		fmt.Fprintf(v, "lowest_52_week_date%s\n", up[0].Lowest_52_week_date)
		fmt.Fprintf(v, "timestamp:%d\n", up[0].Timestamp)

		resp_a, err := http.Get("https://api.upbit.com/v1/orderbook?markets=KRW-BTC")
		if err != nil {
			log.Panicln(err)
		}
		defer resp_a.Body.Close()
		body, err = ioutil.ReadAll(resp_a.Body)
		if err != nil {
			log.Fatal(err)
		}
		var Upbit_Orderbooks []Upbit_Orderbook

		err = json.Unmarshal(body, &Upbit_Orderbooks)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Fprintf(v, "%+v", Upbit_Orderbooks)
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
