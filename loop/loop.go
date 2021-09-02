package main

import (
	"fmt"
	"runtime"
)

func main() {
	// fmt.Println("Hello, World!")

	fmt.Println("Hello, World!", runtime.GOMAXPROCS(-1))

	var (
		i   int
		sum int
	)

	for {
		if i > 5 {
			break
		}
		sum += i
		fmt.Printf("The number %d and sum is %d \n", i, sum)
		i++
	}
}
