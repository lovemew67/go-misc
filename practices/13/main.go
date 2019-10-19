package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// type Value interface {
// 	String() string
// 	Set(string) error
// }

type MySlice []int

func (mySlice *MySlice) String() string {
	return fmt.Sprintf("%v", *mySlice)
}

func (mySlice *MySlice) Set(value string) error {
	if len(*mySlice) > 0 {
		return errors.New("MySlice has already been set!")
	}
	for _, numStr := range strings.Split(value, ",") {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			*mySlice = append(*mySlice, num)
		} else {
			return errors.New("Your input has some non-integers")
		}
	}
	return nil
}

var (
	intPointer    *int
	stringPointer *string
	floatPointer  *float64
	boolPointer   *bool
	intValue      int
	stringValue   string
	floatValue    float64
	boolValue     bool

	mySlice MySlice
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Llongfile)

	intPointer = flag.Int("intPointer", 0, "help message for int pointer")
	stringPointer = flag.String("stringPointer", "", "help message for string pointer")
	floatPointer = flag.Float64("floatPointer", 0.0, "help message for float pointer")
	boolPointer = flag.Bool("boolPointer", false, "help message for bool pointer")

	flag.IntVar(&intValue, "intValue", 0, "help message for int value")
	flag.StringVar(&stringValue, "stringValue", "", "help message for string value")
	flag.Float64Var(&floatValue, "floatValue", 0.0, "help message for float value")
	flag.BoolVar(&boolValue, "boolValue", false, "help message for bool value")

	flag.Var(&mySlice, "mySlice", "a slice of some integers")

	flag.Usage = func() {
		log.Println("**Some other message here**")
		flag.PrintDefaults()
	}
}

func main() {

	flag.Parse()

	log.Printf("NArg: %d", flag.NArg())
	for i, arg := range flag.Args() {
		log.Printf("index: %d, %s & %s\n", i, flag.Arg(i), arg)
	}

	log.Printf("intPointer: %d", *intPointer)
	log.Printf("stringPointer: %s", *stringPointer)
	log.Printf("floatPointer: %f", *floatPointer)
	log.Printf("boolPointer: %t", *boolPointer)

	log.Printf("intValue: %d", intValue)
	log.Printf("stringValue: %s", stringValue)
	log.Printf("floatValue: %f", floatValue)
	log.Printf("boolValue: %t", boolValue)

	log.Printf("mySlice: %+v", mySlice)
}
