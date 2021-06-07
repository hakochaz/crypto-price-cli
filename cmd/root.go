package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypto-price-cli",
	Short: "Get crypto price data directly from the terminal",
	Long:  `Crypto Price is a simple CLI that allows users to get price data - for any cryptocurrency pair`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	priceCmd.PersistentFlags().String("coin", "", "The coin you wish to get the price")
	priceCmd.PersistentFlags().String("vc", "", "The currency to compare to")
}
