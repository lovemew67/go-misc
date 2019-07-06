package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func MarshalJsonAndWriteFile(file interface{}, name string, indent bool) {
	var jsonBytes []byte
	if indent {
		jsonBytes, _ = json.MarshalIndent(file, "", "\t")
	} else {
		jsonBytes, _ = json.Marshal(file)
	}
	err := ioutil.WriteFile(name, jsonBytes, 0644)
	if err != nil {
		log.Printf("ioutil.WriteFile.fail, error: %s\n", err.Error())
	}
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0775)
		if err != nil {
			panic(err)
		}
	}
}
