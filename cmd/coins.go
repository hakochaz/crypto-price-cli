package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"
)

// coinsCmd represents the coins command
var coinsCmd = &cobra.Command{
	Use:   "coins",
	Short: "Get a list of all coins",
	Long:  `Get a list of all compatible coins`,
	Run: func(cmd *cobra.Command, args []string) {
		cl := getCoinsList()
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

func getCoinsList() *types.CoinList {
	cg := gecko.NewClient(nil)

	cl, err := cg.CoinsList()

	if err != nil {
		log.Fatal("Error retrieving coins list")
	}

	return cl
}
