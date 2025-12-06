//- Crie um tipo. O tipo subjacente deve ser int.
//- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
//- Na função main:
//    1. Demonstre o valor da variável "x"
//    2. Demonstre o tipo da variável "x"
//    3. Atribua 42 à variável "x" utilizando o operador "="
//    4. Demonstre o valor da variável "x"
//- Para os aventureiros: https://golang.org/ref/spec#Types
//- Agora já somos quase ninjas nível 1!
//- Solução: https://play.golang.org/p/snm4WuuYmG

package main

import "fmt"

type art int

var x art = 42
var a [32]int

var y string = "James Bond"

func main() {
	fmt.Printf("valor: %v\ntipo: %T\n", x, x)
	fmt.Println(y)
	y = "New James Bond"
	fmt.Println(y)

	fmt.Printf("%v", a[31])
}
