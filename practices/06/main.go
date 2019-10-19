package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/lovemew67/go-misc/practices/06/test"
)

type person struct {
	Name    string  `json:"name1,string"`
	Age     int     `json:"name2,string"`
	Balance float64 `json:"name3,string"`
	Haha    string  `json:"-"`
}

type personEmpty struct {
	Name    string  `json:"name1,omitempty"`
	Age     int     `json:"name2,string"`
	Balance float64 `json:"name3,string"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 06")
}

func main() {

	// json marshal
	p1 := person{
		Name:    "Howard",
		Age:     30,
		Balance: 100.0,
	}
	bytes, errMarshal := json.Marshal(p1)
	if errMarshal != nil {
		log.Panicf("panic, err: %+v", errMarshal)
	}
	log.Println(string(bytes))

	// json marshal, omit empty
	p2 := personEmpty{
		Age:     30,
		Balance: 100.0,
	}
	bytes, errMarshal = json.Marshal(p2)
	if errMarshal != nil {
		log.Panicf("panic, err: %+v", errMarshal)
	}
	log.Println(string(bytes))

	// json unmarshal
	str1 := []byte(`{"name1":"\"Alice\"","name2":"44"}`)
	var p3 person
	errUnmarshal := json.Unmarshal(str1, &p3)
	if errUnmarshal != nil {
		log.Panicf("panic, err: %+v", errUnmarshal)
	}
	log.Printf("%+v", p3)

	// json RawMessage
	// https://blog.dexiang.me/tech/golang-json/
	test.Raw1()

	// json encoder
	// https://gobyexample.com/json
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// json nested struct
	test.Raw2()

	// json marshal indent
	data := map[string]int{
		"a": 1,
		"b": 2,
	}
	json, errMarshalIndent := json.MarshalIndent(data, "OuO: ", "\t")
	if errMarshalIndent != nil {
		log.Fatal(errMarshalIndent)
	}
	log.Println(string(json))
}
