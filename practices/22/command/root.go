package command

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
		},
	}

	sub1Cmd = &cobra.Command{
		Use: "sub1",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub1Cmd Run with args: %v\n", args)
		},
	}

	sub2Cmd = &cobra.Command{
		Use: "sub2",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub2Cmd Run with args: %v\n", args)
		},
	}

	sub3Cmd = &cobra.Command{
		Use: "sub3",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub3Cmd Run with args: %v\n", args)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent Flags
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.go-misc.yaml)")

	// Bind Flags with Config
	_ = viper.BindPFlags(rootCmd.PersistentFlags())

	// Add Subcommand
	sub2Cmd.AddCommand(sub3Cmd)
	sub1Cmd.AddCommand(sub2Cmd)
	rootCmd.AddCommand(sub1Cmd)
}

func initConfig() {
	handlerName := "initConfig"
	config := viper.GetString("config")

	if config != "" {
		fmt.Printf("[%s] config not empty\n", handlerName)
		viper.SetConfigFile(config)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}

		// Search config in home directory with name ".go-misc" (without extension).
		// Default file extension: yaml
		fmt.Printf("[%s] config empty\n", handlerName)
		fmt.Printf("[%s] home path: %s\n", handlerName, home)
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-misc")
	}
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	fmt.Printf("[%s] key: %s\n", handlerName, viper.GetString("key"))
}
