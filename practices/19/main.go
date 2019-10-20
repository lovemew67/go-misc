package main

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 19")
}

func main() {

	// Config
	viper.SetConfigFile("setting.json")
	_ = viper.ReadInConfig()

	// Type assertion
	slice := viper.Get("slice")
	sliceInterface, _ := slice.([]interface{})
	object1, _ := sliceInterface[0].(map[string]interface{})
	object2, _ := sliceInterface[1].(map[string]interface{})

	// Getting Values From Viper
	log.Printf("object1: %+v", object1["key"])
	log.Printf("object2: %+v", object2["key"])

}
