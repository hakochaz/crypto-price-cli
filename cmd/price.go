/*
Copyright © 2021

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	Short: "Get current price data",
	Long:  `Get current price data for any cryptocurrency pair`,
	Run: func(cmd *cobra.Command, args []string) {
		coin, _ := cmd.Flags().GetString("coin")
		vc, _ := cmd.Flags().GetString("vc")
		if coin == "" || vc == "" {
			fmt.Println("vc and coin are required flags.")
		} else {
			pd := getPriceData(coin, vc)
			fmt.Println("Current Price Data")
			fmt.Println("Coin: ", pd.coin)
			fmt.Println("Currency: ", pd.vc)
			fmt.Println("Price: ", pd.price)
		}
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	priceCmd.PersistentFlags().String("coin", "", "The coin you wish to get the price")
	priceCmd.PersistentFlags().String("vc", "", "The currency to compare to")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// priceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type PriceData struct {
	coin  string
	vc    string
	price float32
}

func getPriceData(coin, vc string) PriceData {
	pd := PriceData{}
	cg := gecko.NewClient(nil)
	sp, err := cg.SimplePrice([]string{coin}, []string{vc})
	if err != nil {
		log.Fatal(err)
	}

	c := (*sp)[coin]
	pd.coin = coin
	pd.vc = vc
	pd.price = c[vc]

	return pd
}
