package command

import (
	"log"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use: "sub",
	Short: "sub command",
	Long: `sub command`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("run sub command")
	},
}
