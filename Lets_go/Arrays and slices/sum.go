package sum

// func Sum(n [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += n[i]
// 	}
// 	return sum
// }

func Sum(n []int) (soma int) {
	for _, nu := range n {
		soma += nu
	}
	return
}

// func SumAll(numbersToSum ...[]int) (somaS []int) {
// 	lengthOfNumbers := len(numbersToSum)
// 	somaS = make([]int, lengthOfNumbers)  - 'MAKE' CRIA UM SLICE COM TAMANHO DEFINIDO(COMO SE FOSSE UMA ARRY, mas o tipo continua sendo SLICE)
// 	for i, numbers := range numbersToSum {
// 		somaS[i] = Sum(numbers)
// 	}

// 	return somaS
// }

func SumAll(numbersToSum ...[]int) (somaS []int) {
	for _, numbers := range numbersToSum {
		somaS = append(somaS, Sum(numbers)) // 'APPEND' ADICIONA UM ELEMENTO AO SLICE, SEM PRECISAR DEFINIR O TAMANHO ANTES
	}

	return
}
