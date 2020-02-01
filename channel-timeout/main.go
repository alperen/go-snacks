package main

import (
	"fmt"
	"log"
	"time"
)

func expensiveOperation(timeout time.Duration, timeoutFunc func()) (*int, error) {
	result := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		result <- 1
	}()

	select {
	case res := <-result:
		return &res, nil
	case <-time.After(timeout * time.Second):
		timeoutFunc()
		return nil, fmt.Errorf("timeoutt")
	}

}

func main() {
	log.Println("started")
	a, err := expensiveOperation(time.Duration(3), func() {
		log.Println("timeout fired")
	})

	if err != nil {
		log.Println("err is not nil", err)
	} else {
		log.Println(*a)
	}
}
