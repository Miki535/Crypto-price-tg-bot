package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	tele "gopkg.in/telebot.v4"
)

var Result string

type CoinGeckoResponse struct {
	Bitcoin struct {
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
	Ethereum struct {
		Usd float64 `json:"usd"`
	} `json:"ethereum"`
	Solana struct {
		Usd float64 `json:"usd"`
	} `json:"solana"`
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

	//var lastmessageText string

	var (
		selector = &tele.ReplyMarkup{}

		menu = &tele.ReplyMarkup{ResizeKeyboard: true}

		btnBtc  = menu.Text("Bitcoin")
		btnUsdt = menu.Text("USDT")
		btnEth  = menu.Text("Ethereum")
		btnSol  = menu.Text("Solana")
		btnTon  = menu.Text("Toncoin")

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

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! \nHere you can find the latest cryptocurrency rates for Bitcoin, Ethereum, Solana, USDT, and TON", selector)
	})

	b.Handle("/test", func(c tele.Context) error {
		GetDataFromApi("bitcoin")
		return c.Send(Result)
	})

	b.Handle(&btnPrev, func(c tele.Context) error {
		return c.Send("Choose cryptocurrency:", menu)
	})
	b.Start()
}

func GetDataFromApi(crypto string) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + crypto + "&vs_currencies=usd")
	if err != nil {
		//Add error
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//Add error
	}

	// Розбір JSON-відповіді
	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		//Add error
	}

	Result = fmt.Sprintf("Курс біткоіну на данний момент...$%.2f\n", result.Bitcoin.Usd)
	//		Result := fmt.Sprintf("Курс ефіру на данний момент...$%.2f\n", result.Ethereum.Usd)

}
