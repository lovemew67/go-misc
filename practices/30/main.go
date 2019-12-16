package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

// https://blog.wu-boy.com/2017/03/error-handler-in-golang/
type MyError1 struct {
	Title   string
	Message string
}

func (e MyError1) Error() string {
	return fmt.Sprintf("%v: %v", e.Title, e.Message)
}

type MyError2 struct {
	Title   string
	Message string
}

func (e MyError2) Error() string {
	return fmt.Sprintf("%v: %v", e.Title, e.Message)
}

var (
	ErrTitle1Message1 = errors.New("title1: message1")
)

func IsMyError1(err error) bool {
	_, ok := err.(MyError1)
	return ok
}

func IsMyError2(err error) bool {
	_, ok := err.(MyError2)
	return ok
}

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile | log.Lmicroseconds)
	log.Println("Hello World: 30")
}

func main() {
	before()
	after()
}

func before() {
	err1 := MyError1{"title1", "message1"}
	ok := IsMyError1(err1)
	if !ok {
		log.Fatal("error is not MyError1")
	}
	if err1.Error() != "title1: message1" {
		log.Fatal("message error")
	}

	err2 := MyError2{"title2", "message2"}
	ok = IsMyError2(err2)
	if !ok {
		log.Fatal("error is not MyError2")
	}
	if err2.Error() != "title2: message2" {
		log.Fatal("message error")
	}

	log.Println(reflect.TypeOf(ErrTitle1Message1))
	log.Println(ErrTitle1Message1 == err1)
	log.Println(ErrTitle1Message1.Error() == err1.Error())

	if _, ok := ErrTitle1Message1.(*MyError1); ok {
		log.Println("ErrTitle1Message1 is MyError1")
	} else {
		log.Println("ErrTitle1Message1 is not MyError1")
	}
}

func after() {
	// https://www.flysnow.org/2019/09/06/go1.13-error-wrapping.html
	// https://tonybai.com/2019/10/18/errors-handling-in-go-1-13/
	err := fmt.Errorf("holy fxxk: %w", ErrTitle1Message1)
	if errors.Is(err, ErrTitle1Message1) {
		log.Println("is ErrTitle1Message1")
	} else {
		log.Println("not ErrTitle1Message1")
	}
	log.Println(err)
	log.Printf("%+v", err)

	newErr1 := fmt.Errorf("%w", MyError2{"title2", "message2"})
	newErr2 := fmt.Errorf("%w", newErr1)
	var e1 MyError2
	if errors.As(newErr2, &e1) {
		log.Println("is MyError2")
	} else {
		log.Println("not MyError2")
	}
}
