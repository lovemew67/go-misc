package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/spf13/viper/remote"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 18")
}

func main() {

	// Precedence Order
	// * explicit call to Set
	// * flag
	// * env
	// * config
	// * key/value store
	// * default

	// Establishing Defaults
	viper.SetDefault("top.key", "value_default")

	// Remote Key/Value Store Example - Unencrypted
	_ = viper.AddRemoteProvider("consul", "172.31.35.44:8500", "plaintext")
	viper.SetConfigType("json")
	if err := viper.ReadRemoteConfig(); err != nil {
		log.Panicf("unable to read remote config: %+v", err)
	}

	// Config
	viper.SetConfigFile("setting.json")
	_ = viper.ReadInConfig()

	// Working with Environment Variables
	viper.SetEnvPrefix("spf")
	_ = viper.BindEnv("top.key")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	os.Setenv("SPF_TOP_KEY", "value_env")

	// Working with Flags
	// *Note*, BindPFlag need to cowork with Cobra
	// serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
	// viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	// *Note*, that the variable author will not be set to the value from config (viper), 
	// *Note*, when the --top.key flag is not provided by user.
	sp := pflag.String("top.key", "value_flag_a", "help message for top.key")
	pflag.Lookup("top.key").NoOptDefVal = "value_flag"
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	// Explicit call to Set
	// viper.Set("top.key", "value_set")

	// Getting Values From Viper
	log.Println("top.key:", viper.GetString("top.key"))

	// Retrieve from pflag
	log.Printf("sp: %s", *sp)
}
