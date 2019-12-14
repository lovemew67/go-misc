package main

import (
	"log"

	"github.com/lovemew67/go-misc/practices/29/inner"
	"github.com/lovemew67/go-misc/practices/29/lo"
	"github.com/lovemew67/go-misc/practices/29/lov2"
	"github.com/lovemew67/go-misc/practices/29/yo_v1"
	"github.com/lovemew67/go-misc/practices/29/yo_v2"
)

var (
	strValue = "string"
	strPtr   = &strValue
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile | log.Lmicroseconds)
	log.Println("Hello World: 29")
}

func main() {
	inner.AccessV1()
	inner.AccessV2()

	log.Printf("model v1: %+v", inner.ModelV1{Member: strValue})
	log.Printf("model v2: %+v", inner.ModelV2{Member: strPtr})

	yo_v1.Access()
	yo_v2.Access()

	log.Printf("model yo_v1: %+v", yo_v1.Model{Member: strValue})
	log.Printf("model yo_v2: %+v", yo_v2.Model{Member: strValue})

	lo.Access()
	lov2.Access()

	log.Printf("model lo: %+v", lo.Model{Member: strValue})
	log.Printf("model lov1: %+v", lov2.Model{Member: strValue})
}
