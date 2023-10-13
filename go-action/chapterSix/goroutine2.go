package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg1 sync.WaitGroup

func main333() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg1.Add(2)
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg1.Wait()
	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	defer wg1.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Println("%s:%d\n", prefix, outer)

	}
	fmt.Println("Completed", prefix)
}
