package main

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	// You want -, _, and . in flags to compare the same. aka --my-flag == --my_flag == --my.flag
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 20")
}

func main() {
	flagNormalize()
	envNormalize()
}

func flagNormalize() {
	vFlag := viper.New()

	// Establishing Defaults
	// vFlag.SetDefault("top.key", "value_default")

	// Setting no option default values for flags
	sp := pflag.String("top.key", "1234", "help message")
	pflag.Lookup("top.key").NoOptDefVal = "4321"

	// Mutating or "Normalizing" Flag names
	pflag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	pflag.Parse()
	_ = vFlag.BindPFlags(pflag.CommandLine)

	// Getting Values From Viper
	log.Printf("flag, top.key is set: %t", vFlag.IsSet("top.key"))
	log.Printf("flag, top.key: %s", vFlag.GetString("top.key"))

	// Retrieve from pflag
	log.Printf("sp: %s", *sp)
}

func envNormalize() {
	vEnv := viper.New()

	// Establishing Defaults
	vEnv.SetDefault("top.key", "value_default")

	// Working with Environment Variables
	vEnv.SetEnvPrefix("spf")
	_ = vEnv.BindEnv("top.key")
	_ = vEnv.BindEnv("top.sub.key")
	_ = vEnv.BindEnv("top.sub_key")
	_ = vEnv.BindEnv("top_sub_key")

	// For matching from viper key to env variables, so replace "." with "_"
	vEnv.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set env variables
	os.Setenv("SPF_TOP_KEY", "123")
	os.Setenv("SPF_TOP_SUB_KEY", "456")
	os.Setenv("SPF_ANOTHER_KEY", "789")

	// Getting Values From Viper
	log.Printf("env, top.key: %s", vEnv.GetString("top.key"))
	log.Printf("env, top.sub.key: %s", vEnv.GetString("top.sub.key"))
	log.Printf("env, top.sub_key: %s", vEnv.GetString("top.sub_key"))
	log.Printf("env, top_sub_key: %s", vEnv.GetString("top_sub_key"))

	// Will not be empty when auto
	vEnv.AutomaticEnv()
	log.Printf("env, another.key: %s", vEnv.GetString("another.key"))

	// Allow empty
	_ = vEnv.BindEnv("empty")
	vEnv.AllowEmptyEnv(true)
	os.Setenv("SPF_EMPTY", "")
	log.Printf("env, empty is set: %t", vEnv.IsSet("empty"))
	log.Printf("env, empty: %s", vEnv.GetString("empty"))
}
