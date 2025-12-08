package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return gettingprefix(lang) + name
}

func gettingprefix(lang string) (prefix string) { // O primeiro parentese é o valor de entrada, o segundo é o valor de saída
	//Casso nao tivesse o valor de saída, precisairia especificar a var no return
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// func gettingprefix(lang string) string {
// 	var prefix string
// 	switch lang {
// 	case spanish:
// 		prefix = spanishHelloPrefix
// 	case french:
// 		prefix = frenchHelloPrefix
// 	default:
// 		prefix = englishHelloPrefix
// 	}
// 	return prefix
// }

func main() {
	fmt.Println(Hello("", ""))
}
