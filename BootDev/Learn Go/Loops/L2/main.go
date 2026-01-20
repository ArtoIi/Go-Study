package main

func maxMessages(thresh int) int {
	quantidade := 0
	custo := 0

	for i := 0; ; i++ {
		valor := 100 + i
		if custo+valor > thresh {
			break
		}
		custo += valor
		quantidade++

	}
	return quantidade
}
