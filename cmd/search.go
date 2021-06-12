package cmd

import (
	"fmt"
	"log"

	"github.com/hakochaz/crypto-price-cli/api"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a supported coin by name",
	Long:  `Search for a supported coin by name to get the id for use in price commands`,
	Run: func(cmd *cobra.Command, args []string) {
		n, _ := cmd.Flags().GetString("name")

		if n == "" {
			log.Fatal("Name is a required flag")
		}

		c, err := api.SearchSupportedCoin(n)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Supported Coin Found")
		fmt.Println()
		fmt.Println("Name:", c.Name)
		fmt.Println("ID:", c.Id)
		fmt.Println("Symbol:", c.Symbol)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().String("name", "", "The name of the coin you would like to search for")
}
