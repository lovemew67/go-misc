package command

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "11",
	Short: "Project: 11",
	Long: `Project: 11
a
	b`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("run root command")

		viper.AutomaticEnv()
		viper.SetConfigFile("./setting/setting.toml")
		if errViper := viper.ReadInConfig(); errViper != nil {
			log.Panicln(errViper)
		}

		log.Printf("read config from viper: %s", viper.GetString("app.name"))
	},
}

func Execute() error {
	rootCmd.AddCommand(subCmd)
	return rootCmd.Execute()
}
