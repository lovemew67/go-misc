package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	_ "github.com/spf13/viper/remote"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	yaml = `Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
jacket: leather
trousers: denim
age: 35
eyes : brown
beard: true
`
)

type unmarshal struct {
	K1 string
	K2 string
	K3 string `mapstructure:"k2"`
	k4 string
	k5 string `mapstructure:"k2"`
}

type myFlagA struct{}

func (f myFlagA) HasChanged() bool    { return false }
func (f myFlagA) Name() string        { return "my-flag-A-name" }
func (f myFlagA) ValueString() string { return "my-flag-A-value" }
func (f myFlagA) ValueType() string   { return "string" }

type myFlagB1 struct{}

func (f myFlagB1) HasChanged() bool    { return false }
func (f myFlagB1) Name() string        { return "my-flag-B1-name" }
func (f myFlagB1) ValueString() string { return "my-flag-B1-value" }
func (f myFlagB1) ValueType() string   { return "string" }

type myFlagB2 struct{}

func (f myFlagB2) HasChanged() bool    { return false }
func (f myFlagB2) Name() string        { return "my-flag-B2-name" }
func (f myFlagB2) ValueString() string { return "my-flag-B2-value" }
func (f myFlagB2) ValueType() string   { return "string" }

type myFlagSet struct {
	flags []viper.FlagValue
}

func (f myFlagSet) VisitAll(fn func(viper.FlagValue)) {
	for _, flag := range f.flags {
		fn(flag)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 17")
}

func main() {
	// Working with multiple vipers
	// Establishing Defaults
	v1 := viper.New()
	v2 := viper.New()
	v1.SetDefault("dir", "content")
	v2.SetDefault("dir", "foobar")
	v2.SetDefault("taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	// Getting Values From Viper
	log.Println("v1:", v1.GetString("dir"))
	log.Println("v2:", v2.GetString("dir"))

	if v1.IsSet("dir") { // case-insensitive Setting & Getting
		log.Println("v1 dir enabled")
	}

	log.Printf("v2: %+v", v2)

	// Accessing nested keys
	v3 := viper.New()
	v3.SetConfigFile("setting/sub.toml")
	v3.ReadInConfig()

	// Registering and Using Aliases
	v3.RegisterAlias("top.sub.leaf", "Verbose")                  // case-insensitive
	log.Printf("top.sub.leaf: %t\n", v3.GetBool("top.sub.leaf")) // case-insensitive
	log.Printf("Verbose: %t\n", v3.GetBool("Verbose"))

	// Extract sub-tree
	log.Printf("top.sub: %+v\n", v3.Sub("top.sub"))
	log.Printf("top.sub.tree: %+v\n", v3.Sub("top.sub.tree"))

	// Accessing nested keys
	log.Printf("top.sub.leaf: %t\n", v3.GetBool("top.sub.leaf"))
	log.Printf("top.sub.tree.leaf: %t\n", v3.GetBool("top.sub.tree.leaf"))

	// Marshalling to string
	bs, err := json.MarshalIndent(v3.AllSettings(), "OuO | ", "\t")
	if err != nil {
		log.Fatalf("unable to marshal config to JSON: %+v", err)
	}
	log.Println("v3.AllSettings():", string(bs))

	// Unmarshaling
	// Reading Config Files
	v4 := viper.New()
	v4.SetConfigFile("setting/unmarshal.toml")
	v4.ReadInConfig()
	u := &unmarshal{}
	err = v4.UnmarshalKey("top", &u)
	err = v4.Sub("top").Unmarshal(&u)
	if err != nil {
		log.Fatalf("unable to decode into struct, %+v", err)
	}
	log.Printf("unmarshal: %+v\n", u)

	// Setting Overrides
	// Writing Config Files
	// * WriteConfig
	// * SafeWriteConfig
	// * WriteConfigAs
	// * SafeWriteConfigAs
	v4.Set("key", "value")
	// v4.WriteConfig()                                 // overwrite
	v4.SafeWriteConfig()                                // not overwrite
	v4.WriteConfigAs("setting/new_unmarshal1.toml")     // overwriete
	v4.SafeWriteConfigAs("setting/new_unmarshal2.toml") // not overwrite

	// Watching and re-reading config files
	v4.WatchConfig()
	v4.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
	})
	v4.Set("key2", "value2") // only file change
	// for {
	// }

	// Remote Key/Value Store Example - Unencrypted
	// v5 := viper.New()
	// v5.AddRemoteProvider("consul", "172.31.35.44:8500", "plaintext")
	// v5.SetConfigType("json")
	// if err = v5.ReadRemoteConfig(); err != nil {
	// 	log.Panicf("unable to read remote config: %+v", err)
	// }
	// log.Printf("key: %s", v5.GetString("key"))

	// Watching Changes in Consul - Unencrypted
	// for {
	// 	time.Sleep(time.Second * 5) // delay after each request
	// 	if err := v5.WatchRemoteConfig(); err != nil {
	// 		log.Printf("unable to read remote config: %+v", err)
	// 		continue
	// 	}
	// 	log.Printf("current key: %s", v5.GetString("key"))
	// }

	// Remote Key/Value Store Example - Encrypted
	// v6 := viper.New()
	// v6.AddSecureRemoteProvider("consul", "172.31.35.44:8500", "key", "key/myprivatekey.txt")
	// v6.SetConfigType("json")
	// if err = v6.ReadRemoteConfig(); err != nil {
	// 	log.Panicf("unable to read remote config: %+v", err)
	// }
	// bs, err = json.MarshalIndent(v6.AllSettings(), "OuO | ", "\t")
	// if err != nil {
	// 	log.Fatalf("unable to marshal config to JSON: %+v", err)
	// }
	// log.Println("v6.AllSettings():", string(bs))
	// log.Println("version:", v6.GetString("version"))
	// log.Println("viewmode:", v6.GetString("viewmode"))

	// Reading Config from io.Reader
	v7 := viper.New()
	v7.SetConfigType("yaml")
	yamlBS := []byte(yaml)
	v7.ReadConfig(bytes.NewBuffer(yamlBS))
	log.Printf("yaml name: %s", v7.GetString("name"))

	// Working with Environment Variables
	// * `AutomaticEnv()`
	// * `BindEnv(string...) : error`
	// * `SetEnvPrefix(string)`
	// * `SetEnvKeyReplacer(string...) *strings.Replacer`
	// * `AllowEmptyEnv(bool)`
	v8 := viper.New()
	v8.SetEnvPrefix("spf") // will be uppercased automatically
	v8.BindEnv("id")
	os.Setenv("SPF_ID", "13")
	log.Printf("id: %s", v8.GetString("id"))

	v8.AutomaticEnv()
	// v8.AllowEmptyEnv(false)
	v8.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // replace the first underscore
	os.Setenv("SPF_AAA_BBB", "13")
	log.Printf("AAA_BBB: %s", v8.GetString("aaa_bbb"))

	// Flag interfaces
	v9 := viper.New()
	v9.BindFlagValue("myFlagA", myFlagA{})
	flagSet := myFlagSet{
		flags: []viper.FlagValue{
			myFlagB1{},
			myFlagB2{},
		},
	}
	v9.BindFlagValues(flagSet)
	log.Printf("myFlagA: %+v", v9.Get("myFlagA"))
	log.Printf("my-flag-B1-name: %+v", v9.Get("my-flag-B1-name"))
	log.Printf("my-flag-B2-name: %+v", v9.Get("my-flag-B2-name"))

	// Working with Flags
	// serverCmd.Flags().Int("port", 8999, "Port to run Application server on")
	// viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	_ = flag.Int("intPointer1", 123, "help message for int pointer 1")
	_ = pflag.Int("intPointer2", 456, "help message for int pointer 2")
	_ = pflag.IntP("intPointer3", "i", 789, "help message for int pointer 3")
	pflag.Lookup("intPointer2").NoOptDefVal = "654"
	pflag.Lookup("intPointer3").NoOptDefVal = "987"
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Lookup("intPointer1").NoOptDefVal = "321"
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	log.Printf("intPointer1: %d", viper.GetInt("intPointer1"))
	log.Printf("intPointer2: %d", viper.GetInt("intPointer2"))
	log.Printf("intPointer3: %d", viper.GetInt("intPointer3"))
	log.Printf("intpointer1: %d", viper.GetInt("intpointer1"))
	log.Printf("intpointer2: %d", viper.GetInt("intpointer2"))
	log.Printf("intpointer3: %d", viper.GetInt("intpointer3"))
	log.Printf("only flag's long name, i: %d", viper.GetInt("i"))

	// Precedence Order
	// * explicit call to Set
	// * flag
	// * env
	// * config
	// * key/value store
	// * default
}
