package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	tele "gopkg.in/telebot.v4"
)

var userMessage string
var full_Result string
var CryptoChoose float64

type CoinGeckoResponse struct {
	Bitcoin struct {
		Uah float64 `json:"uah"`
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
	Ethereum struct {
		Uah float64 `json:"uah"`
		Usd float64 `json:"usd"`
	} `json:"ethereum"`
	Solana struct {
		Uah float64 `json:"uah"`
		Usd float64 `json:"usd"`
	} `json:"solana"`
	Tether struct {
		Uah float64 `json:"uah"`
		Usd float64 `json:"usd"`
	} `json:"tether"`
	Dogecoin struct {
		Uah float64 `json:"uah"`
		Usd float64 `json:"usd"`
	} `json:"dogecoin"`
}

func main() {
	pref := tele.Settings{
		Token:  "7434140671:AAGineKwMZ-T6_I0vA92qbcrC0K8A9R7YdU",
		Poller: &tele.LongPoller{Timeout: 1 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		selector     = &tele.ReplyMarkup{}
		uah_selector = &tele.ReplyMarkup{}

		menu = &tele.ReplyMarkup{ResizeKeyboard: true}

		btnBtc  = menu.Text("Bitcoin")
		btnUsdt = menu.Text("Tether")
		btnEth  = menu.Text("Ethereum")
		btnSol  = menu.Text("Solana")
		btnTon  = menu.Text("Dogecoin")

		uah_convert_btn = selector.Data("Convert to UAH ₴", "uahs")
		//Inline buttons
		btnPrev = selector.Data("₿ Cryptocurrency", "crypto")
	)

	menu.Reply(
		menu.Row(btnBtc),
		menu.Row(btnUsdt),
		menu.Row(btnEth),
		menu.Row(btnSol),
		menu.Row(btnTon),
	)

	selector.Inline(
		selector.Row(btnPrev),
	)

	uah_selector.Inline(
		selector.Row(uah_convert_btn),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! \nHere you can find the latest cryptocurrency rates for Bitcoin, Ethereum, Solana, USDT, and TON", selector)
	})

	b.Handle(&btnPrev, func(c tele.Context) error {
		return c.Send("Choose cryptocurrency:", menu)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		userMessage = c.Text()
		GetDataFromApi(userMessage, "usd")
		return c.Send(full_Result, uah_selector)
	})

	b.Handle(&uah_convert_btn, func(c tele.Context) error {
		GetDataFromApi(userMessage, "uah")
		return c.Send(full_Result)
	})

	b.Start()
}

func GetDataFromApi(crypto string, choose string) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + crypto + "&vs_currencies=" + choose)
	if err != nil {
		//Add error
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//Add error
	}

	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		//Add error
	}

	switch crypto {
	case "Bitcoin":
		if choose == "usd" {
			CryptoChoose = result.Bitcoin.Usd
		} else {
			CryptoChoose = result.Bitcoin.Uah
		}
	case "Ethereum":
		if choose == "usd" {
			CryptoChoose = result.Ethereum.Usd
		} else {
			CryptoChoose = result.Ethereum.Uah
		}
	case "Tether":
		if choose == "usd" {
			CryptoChoose = result.Tether.Usd
		} else {
			CryptoChoose = result.Tether.Uah
		}
		CryptoChoose = result.Tether.Usd
	case "Solana":
		if choose == "usd" {
			CryptoChoose = result.Solana.Usd
		} else {
			CryptoChoose = result.Solana.Uah
		}
	case "Dogecoin":
		if choose == "usd" {
			CryptoChoose = result.Dogecoin.Usd
		} else {
			CryptoChoose = result.Dogecoin.Uah
		}
	}
	symbol := "$"
	if choose != "usd" {
		symbol = "₴"
	}

	full_Result = fmt.Sprintf("Курс "+crypto+" на данний момент..."+symbol+"%.2f\n", CryptoChoose)
}
