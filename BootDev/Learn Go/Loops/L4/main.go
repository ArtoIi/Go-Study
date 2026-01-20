package main

import "fmt"

func fizzbuzz() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%v - fizzbuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%v - fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%v - buzz\n", i)
		} else {
			fmt.Printf("%v\n", i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
