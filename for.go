package pauek

import "fmt"

func f1() {
	// _1 OMIT
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	// _1 OMIT
	fmt.Println(sum)
}

func f2() {
	// _2 OMIT
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	// _2 OMIT
	fmt.Println(sum)
}

func f3() {
	// _3 OMIT
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	// _3 OMIT
	fmt.Println(sum)
}

func f4() {
	// _4 OMIT
	for { // forever
	}
	// _4 OMIT
}
