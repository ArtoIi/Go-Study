package main

import "fmt"

type arc string

var a arc = "Texto do tipo arc"

type riders int

var b riders = 101

func main() {

	x := 10
	fmt.Printf("%T", a)

	fmt.Printf("\nb=%v\nx=%v", b, x)

	x = int(b)

	fmt.Printf("\nvalor de x = %v\ntipo de x = %T", x, x)
}
