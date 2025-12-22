package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	msgpro := messages[:]
	msgfree := messages[:2]

	switch plan {
	case planPro:
		return msgpro, nil
	case planFree:
		return msgfree, nil
	default:
		return nil, errors.New("unsupported plan")
	}

}
