package main

import (
	"fmt"
	"log"

	gecko "github.com/enixdark/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	list, err := cg.CoinsList(false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available coins:", len(*list))
}
