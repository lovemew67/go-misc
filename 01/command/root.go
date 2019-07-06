package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	// do necessary initialization
}

var rootCmd = &cobra.Command{
	Use:   "go-misc",
	Short: "go-misc just a place to put some un-important golang code.",
	Long:  `go-misc just a place to put some un-important golang code.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World: 01")
	},
}

func Execute() error {
	rootCmd.AddCommand(NewParseCommand())
	rootCmd.AddCommand(NewSeleniumCommand())
	return rootCmd.Execute()
}
