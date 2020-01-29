// gocui, goroutine 활용
// 업비트에서 매 초 json data를 구조체로 unmarshal 후 출력

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/jroimartin/gocui"
)

var (
	done = make(chan struct{})
	wg   sync.WaitGroup
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

type Upbit_Orderbook struct {
	Market          string           `json:"market"`
	Timestamp       int64            `json:"timestamp"`
	Total_ask_size  float32          `json:"total_ask_size"`
	Total_bid_size  float32          `json:"total_bid_size"`
	Orderbook_units []Orderbook_unit `json:"orderbook_units"`
}

type Polo_Orderbook struct {
	Asks     [][2]interface{} `json:"asks"`
	Bids     [][2]interface{} `json:"bids"`
	IsFrozen string           `json:"isFrozen"`
	Seq      int64            `json:"seq"`
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if k, err := g.SetView("orderbook", maxX/2+1, 1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		k.Wrap = true
	}

	if v, err := g.SetView("main", 1, 1, maxX/2-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true

		/*
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
				fmt.Fprintf(v, "timestamp:%d\n", up[0].Timestamp)   */
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	close(done)
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

	wg.Add(1)
	go counter(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	wg.Wait()
}

func counter(g *gocui.Gui) {
	defer wg.Done()
	var Upbit_Orderbooks []Upbit_Orderbook
	var Upbit_Orderbooks_eth []Upbit_Orderbook
	Polo_Orderbooks := new(Polo_Orderbook)
	i := 0
	for {
		select {
		case <-done:
			return
		case <-time.After(1500 * time.Millisecond):

			g.Update(func(g *gocui.Gui) error {
				v, err := g.View("orderbook")
				if err != nil {
					return err
				}
				v.Clear()
				resp_c, err := http.Get("https://poloniex.com/public?command=returnOrderBook&currencyPair=BTC_ETH&depth=10")
				if err != nil {
					log.Panicln(err)
				}
				defer resp_c.Body.Close()

				body_c, err := ioutil.ReadAll(resp_c.Body)
				if err != nil {
					log.Fatal(err)
				}

				err = json.Unmarshal([]byte(body_c), &Polo_Orderbooks)
				if err != nil {
					fmt.Fprintln(v, "error:", err)
				}

				resp_a, err := http.Get("https://api.upbit.com/v1/orderbook?markets=KRW-BTC")
				if err != nil {
					log.Panicln(err)
				}
				defer resp_a.Body.Close()
				body, err := ioutil.ReadAll(resp_a.Body)
				if err != nil {
					log.Fatal(err)
				}

				err = json.Unmarshal(body, &Upbit_Orderbooks)
				if err != nil {
					fmt.Println("error:", err)
				}

				i++

				fmt.Fprintf(v, "test:%d\n", i)
				fmt.Fprintf(v, "Market:%s\n", Upbit_Orderbooks[0].Market)
				fmt.Fprintf(v, "time:%d\n", Upbit_Orderbooks[0].Timestamp)
				fmt.Fprintf(v, "total_ask_size:%f\n", Upbit_Orderbooks[0].Total_ask_size)
				fmt.Fprintf(v, "total_bid_size:%f\n", Upbit_Orderbooks[0].Total_bid_size)
				fmt.Fprintf(v, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[4].Ask_price)
				fmt.Fprintf(v, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[3].Ask_price)
				fmt.Fprintf(v, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[2].Ask_price)
				fmt.Fprintf(v, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[1].Ask_price)
				fmt.Fprintf(v, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[0].Ask_price)
				fmt.Fprintf(v, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[0].Bid_price)
				fmt.Fprintf(v, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[1].Bid_price)
				fmt.Fprintf(v, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[2].Bid_price)
				fmt.Fprintf(v, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[3].Bid_price)
				fmt.Fprintf(v, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[4].Bid_price)

				m, err := g.View("main")
				if err != nil {
					return err
				}
				m.Clear()
				resp_b, err := http.Get("https://api.upbit.com/v1/orderbook?markets=KRW-ETH")
				if err != nil {
					log.Panicln(err)
				}
				defer resp_b.Body.Close()
				body_m, err := ioutil.ReadAll(resp_b.Body)
				if err != nil {
					log.Fatal(err)
				}
				err = json.Unmarshal(body_m, &Upbit_Orderbooks_eth)
				if err != nil {
					fmt.Println("error:", err)
				}

				fmt.Fprintf(m, "test:%d\n", i)
				fmt.Fprintf(m, "Market:%s\n", Upbit_Orderbooks_eth[0].Market)
				fmt.Fprintf(m, "time:%d\n", Upbit_Orderbooks_eth[0].Timestamp)
				fmt.Fprintf(m, "total_ask_size:%f\n", Upbit_Orderbooks_eth[0].Total_ask_size)
				fmt.Fprintf(m, "total_bid_size:%f\n", Upbit_Orderbooks_eth[0].Total_bid_size)
				fmt.Fprintf(m, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[4].Ask_price)
				fmt.Fprintf(m, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[3].Ask_price)
				fmt.Fprintf(m, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[2].Ask_price)
				fmt.Fprintf(m, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[1].Ask_price)
				fmt.Fprintf(m, "\033[31;1mAsk_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[0].Ask_price)
				fmt.Fprintf(m, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[0].Bid_price)
				fmt.Fprintf(m, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[1].Bid_price)
				fmt.Fprintf(m, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[2].Bid_price)
				fmt.Fprintf(m, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[3].Bid_price)
				fmt.Fprintf(m, "\033[32;1mBid_price:%.0f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[4].Bid_price)

				fmt.Fprintf(m, "\033[32;1mupbit_BTCtoETH:%.6f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[0].Ask_price/Upbit_Orderbooks[0].Orderbook_units[0].Bid_price)
				fmt.Fprintf(m, "\033[32;1mupbit_ETHtoBTC:%.6f\033[0m\n", Upbit_Orderbooks_eth[0].Orderbook_units[0].Bid_price/Upbit_Orderbooks[0].Orderbook_units[0].Ask_price)

				var temp, pi float64
				var temp_1, pi_1 float64
				temp = float64(((10000000 / Upbit_Orderbooks_eth[0].Orderbook_units[0].Ask_price) * 0.9995) - 0.01)
				temp_1 = float64(((10000000 / Upbit_Orderbooks[0].Orderbook_units[0].Ask_price) * 0.9995) - 0.0005)
				fmt.Fprintf(m, "\033[31;1mupbit_btc_value:%.6f\033[0m\n", Upbit_Orderbooks[0].Orderbook_units[0].Ask_price)
				fmt.Fprintf(m, "\033[31;1mupbit_buy_btc:%.6f\033[0m\n", temp_1)
				pi, err = strconv.ParseFloat(Polo_Orderbooks.Asks[0][0].(string), 64)
				if err != nil {
					fmt.Println("error:", err)
				}
				pi_1, err = strconv.ParseFloat(Polo_Orderbooks.Bids[0][0].(string), 64)
				if err != nil {
					fmt.Println("error:", err)
				}
				temp_a := ((temp*pi)*0.9995 - 0.0005) * float64(Upbit_Orderbooks[0].Orderbook_units[0].Bid_price)
				temp_b := ((temp_1/pi_1)*0.9995 - 0.0005) * float64(Upbit_Orderbooks_eth[0].Orderbook_units[0].Bid_price)
				fmt.Fprintf(m, "\033[31;1mpolo_change:%.6f\033[0m\n", temp_b)
				fmt.Fprintf(v, "\033[34;1mETHtoBTC:%.6f\033[0m\n", (temp_a)*0.9995)
				fmt.Fprintf(m, "\033[34;1mBTCtoETH:%.6f\033[0m\n", (temp_b)*0.9995)

				fmt.Fprintf(v, "ETHtoBTC_asks:%s\n", Polo_Orderbooks.Asks[0][0].(string))
				fmt.Fprintf(v, "ETHtoBTC_value:%.6f\n", Polo_Orderbooks.Asks[0][1].(float64))
				fmt.Fprintf(v, "BTCtoETH_bids:%s\n", Polo_Orderbooks.Bids[0][0].(string))
				fmt.Fprintf(v, "BTCtoETH_value:%.6f\n", Polo_Orderbooks.Bids[0][1].(float64))

				temp = float64(Upbit_Orderbooks[0].Orderbook_units[0].Bid_price)
				fmt.Fprintf(m, "\033[351;1mabitis:%.6f\033[0m\n", temp)
				temp = ((temp*0.9995)/float64(Upbit_Orderbooks_eth[0].Orderbook_units[0].Ask_price) - 0.01)
				fmt.Fprintf(m, "\033[351;1mbuyeth:%.6f\033[0m\n", temp)
				temp = (temp * pi) * 0.9995
				fmt.Fprintf(m, "\033[351;1mpolochage:%.6f\033[0m\n", temp)
				temp = temp - 0.0005
				fmt.Fprintf(m, "\033[351;1msendupbit:%.6f\033[0m\n", temp)
				temp = (temp - 1) * float64(Upbit_Orderbooks[0].Orderbook_units[0].Bid_price)
				fmt.Fprintf(m, "\033[351;1mbenefit:%.6f\033[0m\n", temp)

				return nil
			})
		}
	}
}
