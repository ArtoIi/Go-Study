package main

func getMessageCosts(messages []string) []float64 {
	pre := make([]float64, len(messages))
	for i := range messages {
		pre[i] = float64(len(messages[i])) * 0.01
	}

	return pre

}
