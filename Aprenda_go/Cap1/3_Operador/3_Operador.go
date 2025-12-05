package main

// := operador para criar nova variavel (e.g. x := 10)
// * Tipagem automatica
// * funciona apenas dentro de funcoes -> {}

// = operador para atribuir valor a variavel ja criada (e.g. var x int; x = 10)
import "fmt"

func main() {
	a := 10
	b := "ola"
	c := true
	fmt.Printf("valor de a: %v, de tipo: %T\n", a, a)
	fmt.Printf("valor de b: %v, de tipo: %T\n", b, b)
	fmt.Printf("valor de c: %v, de tipo: %T\n", c, c)

}
