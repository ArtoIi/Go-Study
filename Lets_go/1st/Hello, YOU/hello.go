package main

import "fmt"

const hellofix = "Hello, "

func Hello(name string) string {
	return hellofix + name
}

func main() {
	fmt.Println(Hello("Art"))
}
