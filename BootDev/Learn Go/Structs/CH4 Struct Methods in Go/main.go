package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (a authenticationInfo) getBasicAuth() (auto string) {
	auto = fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
	return
}
