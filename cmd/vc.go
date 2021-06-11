package cmd

import (
	"fmt"

	"github.com/hakochaz/crypto-price-cli/api"
	"github.com/spf13/cobra"
)

// vcCmd represents the vc command
var vcCmd = &cobra.Command{
	Use:   "vc",
	Short: "Get a list of supported versus currencies",
	Long:  `List all the supported versus currencies to use when getting prices.`,
	Run: func(cmd *cobra.Command, args []string) {
		cl := api.ListAllSupportedVC()
		fmt.Println("Supported versus currencies: ")
		fmt.Println()
		for _, c := range *cl {
			fmt.Println(c)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(vcCmd)
}
