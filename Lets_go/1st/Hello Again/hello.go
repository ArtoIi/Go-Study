package main

import "fmt"

const hellofix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return hellofix + name
}

func main() {
	fmt.Println(Hello(""))
}
