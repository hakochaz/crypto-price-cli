package cmd

import (
	"fmt"
	"log"

	"github.com/hakochaz/crypto-price-cli/api"
	"github.com/spf13/cobra"
)

// coinsCmd represents the coins command
var coinsCmd = &cobra.Command{
	Use:   "coins",
	Short: "Get a list of all coins",
	Long:  `Get a list of all compatible coins`,
	Run: func(cmd *cobra.Command, args []string) {
		cl, err := api.GetCoinsList()

		if err != nil {
			log.Fatal(err)
		}

		for _, c := range *cl {
			fmt.Println("Name: ", c.Name)
			fmt.Println("ID: ", c.ID)
			fmt.Println("Symbol: ", c.Symbol)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(coinsCmd)
}
