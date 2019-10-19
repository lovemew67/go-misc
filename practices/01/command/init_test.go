package command

import (
	"log"
	"os"
	"testing"
)

func BeforeTest() {
	// do setup
}

func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)
	var p *int
	retCode := 0
	p = &retCode
	BeforeTest()
	defer AfterTest(p)
	*p = m.Run()
}

func AfterTest(ret *int) {
	os.Exit(*ret)
}
