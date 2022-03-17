package main

import (
	"fmt"
	"log"

	gecko "github.com/enixdark/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	tickers, err := cg.CoinsIDTickers("enixdark", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(tickers.Tickers))
	tickers, err = cg.CoinsIDTickers("enixdark", 1)
	fmt.Println(len(tickers.Tickers))
}
