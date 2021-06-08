package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	gecko "github.com/superoo7/go-gecko/v3"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Get current or historical price data",
	Long:  `Get current or historical price data for any cryptocurrency pair`,
	Run: func(cmd *cobra.Command, args []string) {
		coin, _ := cmd.Flags().GetString("coin")
		vc, _ := cmd.Flags().GetString("vc")
		d, _ := cmd.Flags().GetString("date")

		if coin == "" || vc == "" {
			fmt.Println("vc and coin are required flags.")
		} else if d != "" {
			pd := getHistoricalPriceData(coin, vc, d)
			fmt.Println("Coin: ", pd.coin)
			fmt.Println("Currency: ", pd.vc)
			fmt.Println("Price: ", pd.price)
			fmt.Println("Date: ", pd.date)
		} else {
			pd := getCurrentPriceData(coin, vc)
			fmt.Println("Current Price Data")
			fmt.Println("Coin: ", pd.coin)
			fmt.Println("Currency: ", pd.vc)
			fmt.Println("Price: ", pd.price)
		}
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
	priceCmd.PersistentFlags().String("date", "", "The date and time you wish to get data for")
}

type PriceData struct {
	coin  string
	vc    string
	price float64
	date  string
}

// getCurrentPriceData gets current price data for a coin
// versus another currency
func getCurrentPriceData(coin, vc string) PriceData {
	pd := PriceData{}
	cg := gecko.NewClient(nil)

	sp, err := cg.SimpleSinglePrice(coin, vc)
	if err != nil {
		log.Fatal(err)
	}

	c := (*sp)
	pd.coin = coin
	pd.vc = vc
	pd.price = float64(c.MarketPrice)

	return pd
}

// getHistoricalPriceData gets historical price data for a coin
// versus another currency
func getHistoricalPriceData(coin, vc, d string) PriceData {
	pd := PriceData{}
	cg := gecko.NewClient(nil)

	sp, err := cg.CoinsIDHistory(coin, d, true)

	if err != nil {
		log.Fatal(err)
	}

	c := (*sp)
	pd.coin = coin
	pd.vc = vc
	pd.price = c.MarketData.CurrentPrice[vc]
	pd.date = d

	return pd
}
