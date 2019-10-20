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
			fmt.Println(`viper.GetString("test"):`, viper.GetString("test"))
		},
		// TraverseChildren: true,
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
	// rootCmd Flags
	rootCmd.PersistentFlags().StringP("config", "a", "", "config root")
	rootCmd.Flags().StringP("version", "b", "", "version root")

	// sub1Cmd Flags
	sub1Cmd.PersistentFlags().StringP("config", "c", "", "config sub1")
	sub1Cmd.Flags().StringP("version", "b", "", "version sub1")

	// flag default in viper
	sub1Cmd.Flags().StringP("test", "e", "", "test sub1")
	_ = viper.BindPFlag("test", sub1Cmd.Flags().Lookup("test"))
	// viper.SetDefault("test", "NAME HERE")

	// sub2Cmd Flags
	sub2Cmd.PersistentFlags().StringP("version", "d", "", "version sub2")

	// Add Subcommand
	sub2Cmd.AddCommand(sub3Cmd)
	sub1Cmd.AddCommand(sub2Cmd)
	rootCmd.AddCommand(sub1Cmd)

	// rootCmd Flags
	fmt.Println(">>>>>>>>>>>>>>> rootCmd")
	fmt.Printf("rootCmd Flags: \n%s\n", rootCmd.Flags().FlagUsages())
	fmt.Printf("rootCmd LocalFlags: \n%s\n", rootCmd.LocalFlags().FlagUsages())
	fmt.Printf("rootCmd LocalNonPersistentFlags: \n%s\n", rootCmd.LocalNonPersistentFlags().FlagUsages())
	fmt.Printf("rootCmd InheritedFlags: \n%s\n", rootCmd.InheritedFlags().FlagUsages())
	fmt.Printf("rootCmd NonInheritedFlags: \n%s\n", rootCmd.NonInheritedFlags().FlagUsages())
	fmt.Printf("rootCmd PersistentFlags: \n%s\n", rootCmd.PersistentFlags().FlagUsages())

	// sub1Cmd Flags
	fmt.Println(">>>>>>>>>>>>>>> sub1Cmd")
	fmt.Printf("sub1Cmd Flags: \n%s\n", sub1Cmd.Flags().FlagUsages())
	fmt.Printf("sub1Cmd LocalFlags: \n%s\n", sub1Cmd.LocalFlags().FlagUsages())
	fmt.Printf("sub1Cmd LocalNonPersistentFlags: \n%s\n", sub1Cmd.LocalNonPersistentFlags().FlagUsages())
	fmt.Printf("sub1Cmd InheritedFlags: \n%s\n", sub1Cmd.InheritedFlags().FlagUsages())
	fmt.Printf("sub1Cmd NonInheritedFlags: \n%s\n", sub1Cmd.NonInheritedFlags().FlagUsages())
	fmt.Printf("sub1Cmd PersistentFlags: \n%s\n", sub1Cmd.PersistentFlags().FlagUsages())

	// sub2Cmd Flags
	fmt.Println(">>>>>>>>>>>>>>> sub2Cmd")
	fmt.Printf("sub2Cmd Flags: \n%s\n", sub2Cmd.Flags().FlagUsages())
	fmt.Printf("sub2Cmd LocalFlags: \n%s\n", sub2Cmd.LocalFlags().FlagUsages())
	fmt.Printf("sub2Cmd LocalNonPersistentFlags: \n%s\n", sub2Cmd.LocalNonPersistentFlags().FlagUsages())
	fmt.Printf("sub2Cmd InheritedFlags: \n%s\n", sub2Cmd.InheritedFlags().FlagUsages())
	fmt.Printf("sub2Cmd NonInheritedFlags: \n%s\n", sub2Cmd.NonInheritedFlags().FlagUsages())
	fmt.Printf("sub2Cmd PersistentFlags: \n%s\n", sub2Cmd.PersistentFlags().FlagUsages())

	// sub3Cmd Flags
	fmt.Println(">>>>>>>>>>>>>>> sub3Cmd")
	fmt.Printf("sub3Cmd Flags: \n%s\n", sub3Cmd.Flags().FlagUsages())
	fmt.Printf("sub3Cmd LocalFlags: \n%s\n", sub3Cmd.LocalFlags().FlagUsages())
	fmt.Printf("sub3Cmd LocalNonPersistentFlags: \n%s\n", sub3Cmd.LocalNonPersistentFlags().FlagUsages())
	fmt.Printf("sub3Cmd InheritedFlags: \n%s\n", sub3Cmd.InheritedFlags().FlagUsages())
	fmt.Printf("sub3Cmd NonInheritedFlags: \n%s\n", sub3Cmd.NonInheritedFlags().FlagUsages())
	fmt.Printf("sub3Cmd PersistentFlags: \n%s\n", sub3Cmd.PersistentFlags().FlagUsages())
	
	fmt.Println("===============")
}
