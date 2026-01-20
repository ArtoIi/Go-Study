package main

func addEmailsToQueue(emails []string) chan string {

	numEmail := len(emails)
	ch := make(chan string, numEmail)
	return ch
}
