package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msg := [3]string{primary, secondary, tertiary}

	var tentativas [3]int
	sum := 0

	for i, s := range msg {
		sum += len(s)
		tentativas[i] = sum

	}
	return msg, tentativas
}
