package main

import (
	"errors"
	"fmt"
	"log"
)

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return e.Err.Error()
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return e.Name + "not found"
}

var (
	ErrNotFound = errors.New("not found")
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile | log.Lmicroseconds)
	log.Println("Hello World: 30")
}

func main() {
	before()
	after()
}

func before() {
	err1 := errors.New("")
	err2 := fmt.Errorf("%v", err1)

	if e, ok := err2.(*QueryError); ok && e.Err == ErrNotFound {
		log.Println("e.Name wasn't found")
	}
}

func after() {
	err1 := errors.New("")
	err2 := fmt.Errorf("%w", err1)
	log.Printf("%+v", err2)

	//   if err == ErrNotFound { … }
	if errors.Is(err2, ErrNotFound) {
		log.Println("something wasn't found")
	}

	//   if e, ok := err.(*QueryError); ok { … }
	var e *QueryError
	if errors.As(err2, &e) {
		log.Println("err is a *QueryError, and e is set to the error's value")
	}
}
