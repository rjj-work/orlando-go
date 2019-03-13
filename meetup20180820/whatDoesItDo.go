// Problem presented in the Orlando Go meeting 2018.08.20 by Ayan George
// Before you try to run, answer a simple question:
//	What will this program produce ?
package main

import (
	"fmt"
	"time"
)

func printAndLoop(i int) {
	fmt.Printf("printAndLoop(%d)\n", i)
	for {
	}
}

func main() {
	for i := 1; i<= 100; i++ {
		go printAndLoop(i)
	}

	time.Sleep(10*time.Second)
}
