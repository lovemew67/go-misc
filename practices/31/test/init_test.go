package test

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)

	var p *int
	retCode := 0
	p = &retCode

	BeforeTest()
	defer AfterTest(p)

	*p = m.Run()
}

func BeforeTest() {
}

func AfterTest(ret *int) {
	os.Exit(*ret)
}
