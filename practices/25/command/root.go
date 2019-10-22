package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
			fmt.Println(`viper.GetString("version"):`, viper.GetString("version"))
			fmt.Println(`viper.GetString("config"):`, viper.GetString("config"))
			fmt.Println(`viper.GetString("test"):`, viper.GetString("test"))
		},
	}

	sub1Cmd = &cobra.Command{
		Use: "sub1",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub1Cmd Run with args: %v\n", args)
			fmt.Println(`viper.GetString("version"):`, viper.GetString("version"))
			fmt.Println(`viper.GetString("config"):`, viper.GetString("config"))
			fmt.Println(`viper.GetString("test"):`, viper.GetString("test"))
		},
	}

	sub2Cmd = &cobra.Command{
		Use: "sub2",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub2Cmd Run with args: %v\n", args)
			fmt.Println(`viper.GetString("version"):`, viper.GetString("version"))
			fmt.Println(`viper.GetString("config"):`, viper.GetString("config"))
			fmt.Println(`viper.GetString("test"):`, viper.GetString("test"))
		},
	}

	sub3Cmd = &cobra.Command{
		Use: "sub3",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub3Cmd Run with args: %v\n", args)
			fmt.Println(`viper.GetString("version"):`, viper.GetString("version"))
			fmt.Println(`viper.GetString("config"):`, viper.GetString("config"))
			fmt.Println(`viper.GetString("test"):`, viper.GetString("test"))
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// rootCmd Flags
	rootCmd.PersistentFlags().StringP("config", "a", "config root", "config root")
	rootCmd.Flags().StringP("version", "b", "version root", "version root")

	// sub1Cmd Flags
	sub1Cmd.PersistentFlags().StringP("config", "c", "config sub1", "config sub1")
	sub1Cmd.Flags().StringP("version", "b", "version sub1", "version sub1")

	// sub2Cmd Flags
	sub2Cmd.PersistentFlags().StringP("version", "d", "version sub2", "version sub2")

	// can bind only default values
	_ = viper.BindPFlags(rootCmd.PersistentFlags())
	_ = viper.BindPFlags(rootCmd.Flags())
	_ = viper.BindPFlags(sub1Cmd.PersistentFlags())
	_ = viper.BindPFlags(sub1Cmd.Flags())
	_ = viper.BindPFlags(sub2Cmd.PersistentFlags())
	_ = viper.BindPFlags(sub2Cmd.Flags())

	// can lookup dynamicly
	_ = viper.BindPFlag("version", sub1Cmd.Flags().Lookup("version"))
	// _ = viper.BindPFlag("test", sub1Cmd.Flags().Lookup("version"))
	// _ = viper.BindPFlag("test", sub2Cmd.Flags().Lookup("version"))

	// Add Subcommand
	sub2Cmd.AddCommand(sub3Cmd)
	sub1Cmd.AddCommand(sub3Cmd)
	sub1Cmd.AddCommand(sub2Cmd)
	rootCmd.AddCommand(sub3Cmd)
	rootCmd.AddCommand(sub1Cmd)
}
