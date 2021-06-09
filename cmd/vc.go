package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"
)

// vcCmd represents the vc command
var vcCmd = &cobra.Command{
	Use:   "vc",
	Short: "Get a list of supported versus currencies",
	Long:  `List all the supported versus currencies to use when getting prices.`,
	Run: func(cmd *cobra.Command, args []string) {
		cl := listAllSupportedVC()
		fmt.Println("Supported versus currencies: ")
		for _, c := range *cl {
			fmt.Println(c)
		}
	},
}

func init() {
	rootCmd.AddCommand(vcCmd)
}

func listAllSupportedVC() *types.SimpleSupportedVSCurrencies {
	cg := gecko.NewClient(nil)
	cl, err := cg.SimpleSupportedVSCurrencies()

	if err != nil {
		log.Fatal("Error retrieving supported currencies list")
	}

	return cl
}
