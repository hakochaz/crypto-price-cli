package api

import (
	"errors"
	"fmt"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"
)

func GetCoinsList() (*types.CoinList, error) {
	cg := gecko.NewClient(nil)

	cl, err := cg.CoinsList()

	if err != nil {
		return nil, err
	}

	return cl, nil
}

type PriceData struct {
	Coin   string
	VC     string
	Price  float64
	Date   string
	Amount float64
	Value  float64
}

// getCurrentPriceData gets current price data for a coin
// versus another currency
func GetCurrentPriceData(id, vc string) PriceData {
	pd := PriceData{}
	cg := gecko.NewClient(nil)

	sp, err := cg.SimpleSinglePrice(id, vc)
	if err != nil {
		log.Fatal(err)
	}

	c := (*sp)
	pd.Coin = id
	pd.VC = vc
	pd.Price = float64(c.MarketPrice)

	return pd
}

// GetHistoricalPriceData gets historical price data for a coin
// versus another currency
func GetHistoricalPriceData(id, vc, d string) PriceData {
	pd := PriceData{}
	cg := gecko.NewClient(nil)

	sp, err := cg.CoinsIDHistory(id, d, true)

	if err != nil {
		log.Fatal(err)
	}

	c := (*sp)
	pd.Coin = id
	pd.VC = vc
	pd.Price = c.MarketData.CurrentPrice[vc]
	pd.Date = d

	return pd
}

// calculateAmount will calculate the value fo the specified
// amount of crypto
func (pd *PriceData) CalculateAmount() {
	pd.Value = pd.Amount * pd.Price
}

// PrintPriceData prints to the console all the related price
// data from the PriceData type
func (pd *PriceData) PrintPriceData() {
	if pd.Date != "" {
		fmt.Println("Historical Price Data")
	} else {
		fmt.Println("Current Price Data")
	}

	fmt.Println()
	fmt.Println("Coin: ", pd.Coin)
	fmt.Println("Currency: ", pd.VC)
	fmt.Printf("Price: %.2f\n", pd.Price)

	if pd.Amount != 0 {
		fmt.Println("Amount: ", pd.Amount)
		fmt.Printf("Value: %.2f\n", pd.Value)
	}
	fmt.Println()
}

type Coin struct {
	Name   string
	Id     string
	Symbol string
}

// SearchCoinFromList takes in a name string and returns
// a Coin type and error
func SearchCoinFromList(name string) (Coin, error) {
	rc := Coin{}
	cl, err := GetCoinsList()

	if err != nil {
		return rc, err
	}

	for _, c := range *cl {
		if c.Name == name || c.ID == name || c.Symbol == name {
			rc = Coin{c.Name, c.ID, c.Symbol}
			return rc, nil
		}
	}

	return rc, errors.New("unable to find supported coin")
}

// ListAllSupportedVC gets all the supproted versus currencies
// for the CoinGecko API
func ListAllSupportedVC() *types.SimpleSupportedVSCurrencies {
	cg := gecko.NewClient(nil)
	cl, err := cg.SimpleSupportedVSCurrencies()

	if err != nil {
		log.Fatal("Error retrieving supported currencies list")
	}

	return cl
}
