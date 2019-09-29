package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

const (
	tagName = "validate"
)

var (
	mailReg = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
)

type User struct {
	Id    int    `validate:"number,min=1,max=1000"`
	Name  string `validate:"string,min=2,max=10"`
	Bio   string `validate:"string"`
	Email string `validate:"email"`
}

type Validator interface {
	Validate(interface{}) (bool, error)
}

type DefaultValidator struct {}
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

type StringValidator struct {
	Min int
	Max int
}
func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))
	if l == 0 {
	  	return false, fmt.Errorf("cannot be blank")
	}
	if l < v.Min {
		return false, fmt.Errorf("should be at least %v chars long", v.Min)
	}
	if v.Max >= v.Min && l > v.Max {
	  	return false, fmt.Errorf("should be less than %v chars long", v.Max)
	}
	return true, nil
}

type NumberValidator struct {
	Min int
	Max int
}
func (v NumberValidator) Validate(val interface{}) (bool, error) {
	num := val.(int)
	if num < v.Min {
	  	return false, fmt.Errorf("should be greater than %v", v.Min)
	}
	if v.Max >= v.Min && num > v.Max {
	  	return false, fmt.Errorf("should be less than %v", v.Max)
	}
	return true, nil
}

type EmailValidator struct {}
func (v EmailValidator) Validate(val interface{}) (bool, error) {
  	if !mailReg.MatchString(val.(string)) {
    	return false, fmt.Errorf("is not a valid email address")
  	}
  	return true, nil
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")
	switch args[0] {
	case "number":
	  	validator := NumberValidator{}
	  	fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
	  	return validator
	case "string":
	  	validator := StringValidator{}
	  	fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
	  	return validator
	case "email":
	  	return EmailValidator{}
	}
	return DefaultValidator{}
}

func validateStruct(s interface{}) []error {
	errs := []error{}

	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {

		// Get the field, returns https://golang.org/pkg/reflect/#StructField
	  	tag := v.Type().Field(i).Tag.Get(tagName)

	  	// Skip if tag is not defined or ignored
	  	if tag == "" || tag == "-" {
			continue
		}
		  
	  	// Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)
		  
	  	// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())
		  
	  	// Append error to results
	  	if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s: %s", v.Type().Field(i).Name, err.Error()))
	  	}
	}
	return errs
}  

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 07")
}

// https://medium.com/swlh/9a7aeedcdc5b
func main() {
	user := User{
		Id:    0,
		Name:  "superlongstring",
		Bio:   "",
		Email: "foobar",
	}
	log.Println("Errors:")
	for i, err := range validateStruct(user) {
		log.Printf("\t%d. %s\n", i+1, err.Error())
	}
}
