package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func work(process int, wg *sync.WaitGroup, res *int) {
	defer wg.Done()

	r := rand.Intn(4) + 1
	log.Printf("input: %d proc will work %d seconds\n", process, r)
	time.Sleep(time.Duration(r) * time.Second)
	*res = process * 2
}

func main() {
	var wg sync.WaitGroup
	processes := []int{2, 3, 5, 6, 8, 9}
	results := make([]int, len(processes))

	wg.Add(len(processes))

	startTime := time.Now()
	log.Printf("%s", startTime)

	for index, process := range processes {
		go work(process, &wg, &results[index])
	}

	wg.Wait()
	log.Println("All processes have done")
	log.Println("Execution takes the longest goroutine alone.", time.Since(startTime))
	log.Println(results)
}
