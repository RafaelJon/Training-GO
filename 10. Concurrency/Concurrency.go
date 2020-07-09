package main

import (
	"fmt"
	"sync"
	"time"
)

func recoverFibo(wg *sync.WaitGroup) {
	if r := recover(); r != nil {
		fmt.Println("Recovered from", r)
	}
	wg.Done()
}

func printFibonaci(number int, wg *sync.WaitGroup) {
	defer recoverFibo(wg)
	fibo := 1
	prevFibo := 1
	for i := 0; i < number; i++ {
		temp := fibo
		fibo = fibo + prevFibo
		prevFibo = temp
		if i == 5 {
			panic("error: fibo failed")
		}
		time.Sleep(time.Second)
	}
	fmt.Println(fibo)

}

// func FIBO(number int) {
// 	fibo := 1
// 	prevFibo := 1
// 	for i := 0; i < number; i++ {
// 		temp := fibo
// 		fibo = fibo + prevFibo
// 		prevFibo = temp
// 		time.Sleep(time.Second)
// 	}
// 	fmt.Println(fibo)
// }

func main() {
	defer fmt.Println("main function finished")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printFibonaci(i, &wg)
	}

	// for i := 0; i < 100; i++ {
	// 	FIBO(i)
	// }

	wg.Wait()
}
