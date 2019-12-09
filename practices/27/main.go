package main

import (
	"context"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds | log.Lshortfile)
	log.Println("Hello World: 27")
}

func main() {
	rootCtx := context.Background()
	log.Printf("%t", rootCtx.Done() == nil)

	lv1CancelCtx, lv1CancelFunc := context.WithCancel(rootCtx)
	lv2CancelCtx, lv2CancelFunc := context.WithCancel(lv1CancelCtx)
	select {
	case <-lv1CancelCtx.Done():
		log.Println("lv1cancelCtx cancelled")
	default:
		log.Println("lv1cancelCtx not cancelled")
	}
	select {
	case <-lv2CancelCtx.Done():
		log.Println("lv2CancelCtx cancelled")
	default:
		log.Println("lv2CancelCtx not cancelled")
	}

	lv1CancelFunc()
	select {
	case <-lv1CancelCtx.Done():
		log.Println("lv1cancelCtx cancelled")
	default:
		log.Println("lv1cancelCtx not cancelled")
	}
	select {
	case <-lv2CancelCtx.Done():
		log.Println("lv2CancelCtx cancelled")
	default:
		log.Println("lv2CancelCtx not cancelled")
	}

	lv2CancelFunc()
	select {
	case <-lv1CancelCtx.Done():
		log.Println("lv1cancelCtx cancelled")
	default:
		log.Println("lv1cancelCtx not cancelled")
	}
	select {
	case <-lv2CancelCtx.Done():
		log.Println("lv2CancelCtx cancelled")
	default:
		log.Println("lv2CancelCtx not cancelled")
	}

	lv3CancelCtx, _ := context.WithCancel(lv2CancelCtx)
	select {
	case <-lv3CancelCtx.Done():
		log.Println("lv3CancelCtx cancelled")
	default:
		log.Println("lv3CancelCtx not cancelled")
	}

	lv1DeadlineCtx, _ := context.WithDeadline(rootCtx, time.Now().Add(time.Duration(100)*time.Millisecond))
	<-lv1DeadlineCtx.Done()
	log.Println("lv1DeadlineCtx cancelled")
	log.Printf("lv1DeadlineCtx error: %+v\n", lv1DeadlineCtx.Err())

	lv1TimeoutCtx, _ := context.WithTimeout(rootCtx, time.Duration(100)*time.Millisecond)
	<-lv1TimeoutCtx.Done()
	log.Println("lv1TimeoutCtx cancelled")
	log.Printf("lv1TimeoutCtx error: %+v\n", lv1TimeoutCtx.Err())

	lv2Cancel2Ctx, _ := context.WithTimeout(lv1TimeoutCtx, time.Duration(100)*time.Millisecond)
	<-lv2Cancel2Ctx.Done()
	log.Println("lv2Cancel2Ctx cancelled")
	log.Printf("lv2Cancel2Ctx error: %+v\n", lv2Cancel2Ctx.Err())

	lv1ValueCtx := context.WithValue(rootCtx, "key", "value")
	if lv1ValueCtx.Done() == nil {
		log.Println("parent ctx is rootCtx")
	}

	lv1CancelCtx2, lv1CancelFunc2 := context.WithCancel(rootCtx)
	lv2ValueCtx := context.WithValue(lv1CancelCtx2, "key", "value")
	select {
	case <-lv2ValueCtx.Done():
		log.Println("lv2ValueCtx cancelled")
	default:
		log.Println("lv2ValueCtx not cancelled")
	}
	lv1CancelFunc2()
	select {
	case <-lv2ValueCtx.Done():
		log.Println("lv2ValueCtx cancelled")
	default:
		log.Println("lv2ValueCtx not cancelled")
	}

	log.Println(lv1CancelCtx2.Value("key"))
	log.Println(lv2ValueCtx.Value("key"))

}
