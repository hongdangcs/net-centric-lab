package lab2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func SimulateLibrary(studentCount int, seatCount int) {

	seats := make(chan int, seatCount)

	var wg sync.WaitGroup
	wg.Add(studentCount)

	timeZero := time.Now().Second()
	for i := 0; i < studentCount; i++ {
		go func(i int) {
			defer wg.Done()
			arrivalTime := time.Duration(rand.Intn(5)) * time.Second
			time.Sleep(arrivalTime)
			if len(seats) == cap(seats) {
				fmt.Printf("Time %d: Student %d is waiting\n", time.Now().Second()-timeZero, i)
			}
			seats <- i
			fmt.Printf("Time %d: Student %d start reading at the lib\n", time.Now().Second()-timeZero, i)
			readingTime := time.Duration(rand.Intn(4)+1) * time.Second
			time.Sleep(readingTime)

			<-seats

			fmt.Printf("Time %d: Student %d is leaving. Spent %d hours reading\n", time.Now().Second()-timeZero, i, readingTime/time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Time %d: No more students. Let's call it a day\n", time.Now().Second()-timeZero)
}
