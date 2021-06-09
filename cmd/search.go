package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a coin by name",
	Long:  `Search for a coin by name to get the id to use in price commands`,
	Run: func(cmd *cobra.Command, args []string) {
		n, _ := cmd.Flags().GetString("name")

		if n == "" {
			log.Fatal("Name is a required flag")
		}

		c, err := searchCoinFromList(n)

		if err != nil {
			log.Fatal("Unable to find supported coin")
		}

		fmt.Println("Supported Coin Found")
		fmt.Println()
		fmt.Println("Name:", c.name)
		fmt.Println("ID:", c.id)
		fmt.Println("Symbol:", c.symbol)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().String("name", "", "The name of the coin you would like to search for")
}

type Coin struct {
	name   string
	id     string
	symbol string
}

func searchCoinFromList(name string) (Coin, error) {
	cl := getCoinsList()
	rc := Coin{}

	for _, c := range *cl {
		if c.Name == name || c.ID == name || c.Symbol == name {
			rc = Coin{c.Name, c.ID, c.Symbol}
			return rc, nil
		}
	}

	return rc, errors.New("unable to find supported coin")
}