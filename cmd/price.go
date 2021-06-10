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
		id, _ := cmd.Flags().GetString("id")
		vc, _ := cmd.Flags().GetString("vc")
		d, _ := cmd.Flags().GetString("date")
		a, _ := cmd.Flags().GetFloat64("amount")

		pd := PriceData{}

		if id == "" || vc == "" {
			fmt.Println("vc and id are required flags for the price command")
		} else if d != "" {
			pd = getHistoricalPriceData(id, vc, d)
		} else {
			pd = getCurrentPriceData(id, vc)
		}

		if a > 0 {
			pd.amount = a
			pd.calculateAmount()
		}

		printPriceData(pd)
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
	priceCmd.PersistentFlags().String("id", "", "The identifier of the coin you wish to get the price")
	priceCmd.PersistentFlags().String("vc", "", "The currency to compare versus")
	priceCmd.PersistentFlags().String("date", "", "The date and time you wish to get data for")
	priceCmd.PersistentFlags().Float64("amount", 0, "Get price data for a specified amount of the coin")
}

type PriceData struct {
	coin   string
	vc     string
	price  float64
	date   string
	amount float64
	value  float64
}

// getCurrentPriceData gets current price data for a coin
// versus another currency
func getCurrentPriceData(id, vc string) PriceData {
	pd := PriceData{}
	cg := gecko.NewClient(nil)

	sp, err := cg.SimpleSinglePrice(id, vc)
	if err != nil {
		log.Fatal(err)
	}

	c := (*sp)
	pd.coin = id
	pd.vc = vc
	pd.price = float64(c.MarketPrice)

	return pd
}

// getHistoricalPriceData gets historical price data for a coin
// versus another currency
func getHistoricalPriceData(id, vc, d string) PriceData {
	pd := PriceData{}
	cg := gecko.NewClient(nil)

	sp, err := cg.CoinsIDHistory(id, d, true)

	if err != nil {
		log.Fatal(err)
	}

	c := (*sp)
	pd.coin = c.Name
	pd.vc = vc
	pd.price = c.MarketData.CurrentPrice[vc]
	pd.date = d

	return pd
}

// calculateAmount will calculate the value fo the specified
// amount of crypto
func (pd *PriceData) calculateAmount() {
	pd.value = pd.amount * pd.price
}

// printPriceData prints to the console all the related price
// data based on the commands flags
func printPriceData(pd PriceData) {
	if pd.date != "" {
		fmt.Println("Historical Price Data")
	} else {
		fmt.Println("Current Price Data")
	}

	fmt.Println("Coin: ", pd.coin)
	fmt.Println("Currency: ", pd.vc)
	fmt.Printf("Price: %.2f\n", pd.price)

	if pd.amount != 0 {
		fmt.Println("Amount: ", pd.amount)
		fmt.Printf("Value: %.2f\n", pd.value)
	}
}
