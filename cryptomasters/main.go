package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"frontendmasters.com/cryptomasters/currencyapi"
)

func main() {
	if len(os.Args) > 1 {
		var wg sync.WaitGroup
		for _, arg := range os.Args[1:] {
			wg.Add(1)
			go func(currency string) {
				defer wg.Done()
				printRate(currency)
			}(arg)
		}
		wg.Wait()
	} else {
		var currencyCode string
		for len(currencyCode) == 0 {
			fmt.Print("Enter the currency code: ")
			fmt.Scan(&currencyCode)
			printRate(currencyCode)
		}
	}
}

func printRate(currency string) {
	fmt.Printf("Fetching the price for %v ⌛️\n", currency)
	price, err := currencyapi.GetRate(strings.ToUpper(currency))

	if err != nil {
		fmt.Println("We coudn't get the rate. Details:")
		fmt.Println(err)
	} else {
		fmt.Printf("The current price of %v is %.8f \n", price.Currency, price.Price)
	}
}
