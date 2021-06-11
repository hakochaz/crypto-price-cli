package cmd

import (
	"fmt"

	"github.com/hakochaz/crypto-price-cli/api"
	"github.com/spf13/cobra"
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

		pd := api.PriceData{}

		if id == "" || vc == "" {
			fmt.Println("vc and id are required flags for the price command")
		} else if d != "" {
			pd = api.GetHistoricalPriceData(id, vc, d)
		} else {
			pd = api.GetCurrentPriceData(id, vc)
		}

		if a > 0 {
			pd.Amount = a
			pd.CalculateAmount()
		}

		pd.PrintPriceData()
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
	priceCmd.PersistentFlags().String("id", "", "The identifier of the coin you wish to get the price")
	priceCmd.PersistentFlags().String("vc", "", "The currency to compare versus")
	priceCmd.PersistentFlags().String("date", "", "The date and time you wish to get data for")
	priceCmd.PersistentFlags().Float64("amount", 0, "Get price data for a specified amount of the coin")
}
