package command

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Assign flags to a command
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use: "21 [commands]",
		// There is not context between Pre/Post Run hooks
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd Run with args: %v\n", args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
		},
	}

	sub1Cmd = &cobra.Command{
		Use:   "sub1",
		Short: "sub1",
		Long:  `sub1`,
		Args:  cobra.MinimumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub1Cmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub1Cmd Run with args: %v\n", args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub1Cmd PostRun with args: %v\n", args)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub1Cmd PersistentPostRun with args: %v\n", args)
		},
		// Local Flag on Parent Commands
		TraverseChildren: true,
	}

	sub2Cmd = &cobra.Command{
		Use:   "sub2",
		Short: "sub2",
		Long:  `sub2`,
		Args:  cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub2Cmd PersistentPreRun with args: %v\n", args)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub2Cmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub2Cmd Run with args: %v\n", args)
		},
		// Local Flag on Parent Commands
		TraverseChildren: true,
	}

	sub3Cmd = &cobra.Command{
		Use:   "sub3",
		Short: "sub3",
		Long:  `sub3`,
		Args:  cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub3Cmd PreRun with args: %v\n", args)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside sub3Cmd Run with args: %v\n", args)
		},
		// Local Flag on Parent Commands
		TraverseChildren: true,
	}
)

func Execute() error {
	// For testing
	fmt.Println(`>>>>>>>>>> rootCmd.SetArgs([]string{""})`)
	rootCmd.SetArgs([]string{""})
	_ = rootCmd.Execute()
	fmt.Println()

	// For testing
	fmt.Println(`>>>>>>>>>> rootCmd.SetArgs([]string{"arg1", "arg2"})`)
	rootCmd.SetArgs([]string{"arg1", "arg2"})
	_ = rootCmd.Execute()
	fmt.Println()

	// For testing
	fmt.Println(`>>>>>>>>>> rootCmd.SetArgs([]string{"sub1", "arg1", "arg2"})`)
	rootCmd.SetArgs([]string{"sub1", "arg1", "arg2"})
	_ = rootCmd.Execute()
	fmt.Println()

	// For testing
	fmt.Println(`>>>>>>>>>> rootCmd.SetArgs([]string{"sub1", "sub2"})`)
	rootCmd.SetArgs([]string{"sub1", "sub2"})
	_ = rootCmd.Execute()
	fmt.Println()

	// For testing
	fmt.Println(`>>>>>>>>>> rootCmd.SetArgs([]string{"sub1", "sub2", "arg1"})`)
	rootCmd.SetArgs([]string{"sub1", "sub2", "arg1"})
	_ = rootCmd.Execute()
	fmt.Println()

	// For testing
	fmt.Println(`>>>>>>>>>> rootCmd.SetArgs([]string{"sub1", "sub2", "sub3"})`)
	rootCmd.SetArgs([]string{"sub1", "sub2", "sub3"})
	_ = rootCmd.Execute()
	fmt.Println()

	// For testing
	fmt.Println(`>>>>>>>>>> rootCmd.SetArgs([]string{"sub1", "sub2", "sub3", "arg1"})`)
	rootCmd.SetArgs([]string{"sub1", "sub2", "sub3", "arg1"})
	_ = rootCmd.Execute()
	fmt.Println()

	// Normal execution
	fmt.Println(`>>>>>>>>>> rootCmd.SetArgs(os.Args[1:])`)
	rootCmd.SetArgs(os.Args[1:])
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent Flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")

	// Local Flags
	// Version Flag
	// The template can be customized using the `cmd.SetVersionTemplate(s string)` function.
	rootCmd.Flags().StringP("version", "v", "", "display the version")

	// Required flags
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	_ = rootCmd.MarkFlagRequired("viper") // work on local flag

	// Bind Flags with Config
	// 	**Note**, that the variable `author` will not be set to the value from config,
	//  **Note**, when the `--author` flag is not provided by user.
	_ = viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	_ = viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE lovemew67@gmail.com")
	viper.SetDefault("license", "apache")

	// Add Subcommand
	sub2Cmd.AddCommand(sub3Cmd)
	sub1Cmd.AddCommand(sub2Cmd)
	rootCmd.AddCommand(sub1Cmd)

	// Other about handy usage
	// ### Help Message
	// ```go
	// cmd.SetHelpCommand(cmd *Command)
	// cmd.SetHelpFunc(f func(*Command, []string))
	// cmd.SetHelpTemplate(s string)
	// ```
	// ## Usage Message
	// ```go
	// cmd.SetUsageFunc(f func(*Command) error)
	// cmd.SetUsageTemplate(s string)
	// ```
	// ## Suggestions when "unknown command" happens
	// ## Generating documentation for your command
	// ## Generating bash completions
	// ## Generating zsh completions
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
