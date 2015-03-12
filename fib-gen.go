// +build OMIT

package main

import "fmt"

func fib(c chan int, n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		c <- a // HL
	}
	close(c) // HL
}

func main() {
	c := make(chan int) // HL
	go fib(c, 10)       // HL

	for x := range c {
		fmt.Println(x)
	}
}
