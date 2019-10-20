package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	goflag "flag"

	"github.com/spf13/pflag"
)

// type Value interface {
// 	String() string
// 	Set(string) error
// }

// type Value interface {
// 	String() string
// 	Set(string) error
// 	Type() string
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

func (mySlice *MySlice) Type() string {
	return "mySlice"
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

	intPointer2  *int
	boolPointer2 *bool
	intValue2    int
	boolValue2   bool

	mySlice2 MySlice

	ip  *int
	ip2 *int
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

func aliasNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	// You want to alias two flags. aka --old-flag-name == --new-flag-name
	switch name {
	case "old-flag-name":
		name = "new-flag-name"
		break
	}
	return pflag.NormalizedName(name)
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Llongfile)
	log.Println("Hello World: 14")

	intPointer = pflag.Int("intPointer", 0, "help message for int pointer")
	stringPointer = pflag.String("stringPointer", "", "help message for string pointer")
	floatPointer = pflag.Float64("floatPointer", 0.0, "help message for float pointer")
	boolPointer = pflag.Bool("boolPointer", false, "help message for bool pointer")

	pflag.IntVar(&intValue, "intValue", 0, "help message for int value")
	pflag.StringVar(&stringValue, "stringValue", "", "help message for string value")
	pflag.Float64Var(&floatValue, "floatValue", 0.0, "help message for float value")
	pflag.BoolVar(&boolValue, "boolValue", false, "help message for bool value")

	pflag.Var(&mySlice, "mySlice", "a slice of some integers")

	pflag.Usage = func() {
		log.Println("**Some other message here**")
		pflag.PrintDefaults()
	}

	intPointer2 = pflag.IntP("intPointer2", "a", 0, "help message for int pointer 2")
	boolPointer2 = pflag.BoolP("boolPointer2", "b", false, "help message for bool pointer 2")
	pflag.IntVarP(&intValue2, "intValue2", "d", 0, "help message for int value 2")
	pflag.BoolVarP(&boolValue2, "boolValue2", "e", false, "help message for bool value 2")
	pflag.VarP(&mySlice2, "mySlice3", "f", "a slice of some integers")

	// Setting no option default values for flags
	ip = pflag.IntP("flagname", "g", 1234, "help message")
	pflag.Lookup("flagname").NoOptDefVal = "4321"

	// Mutating or "Normalizing" Flag names
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)
	pflag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	// Deprecating a flag or its shorthand
	pflag.CommandLine.MarkDeprecated("badflag", "please use --good-flag instead")
	pflag.CommandLine.MarkShorthandDeprecated("noshorthandflag", "please use --noshorthandflag only")

	// Hidden flags
	pflag.CommandLine.MarkHidden("secretFlag")

	// Supporting Go flags when using pflag
	ip2 = goflag.Int("newflagname", 1234, "help message for newflagname")
}

func main() {

	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()

	log.Printf("NArg: %d", pflag.NArg())
	for i, arg := range pflag.Args() {
		log.Printf("index: %d, %s & %s\n", i, pflag.Arg(i), arg)
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

	log.Printf("intPointer2: %d", *intPointer2)
	log.Printf("boolPointer2: %t", *boolPointer2)

	log.Printf("intValue2: %d", intValue2)
	log.Printf("boolValue2: %t", boolValue2)

	log.Printf("mySlice3: %+v", mySlice2)

	// |Parsed Arguments|Resulting Value|
	// |–flagname=1357  |ip=1357        |
	// |–flagname	     |ip=4321        |
	// |[nothing]	     |ip=1234        |

	// Disable sorting of flags
	log.Println("before disable sorting")
	pflag.CommandLine.PrintDefaults()
	pflag.CommandLine.SortFlags = false
	log.Println("after disable sorting")
	pflag.CommandLine.PrintDefaults()
}
