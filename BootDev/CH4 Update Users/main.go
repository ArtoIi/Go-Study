package main

type User struct {
	Name string
	membership
}
type membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) (usuario User) {
	if membershipType == "premium" {
		usuario.Type = "premium"
		usuario.Name = name
		usuario.MessageCharLimit = 1000
		return
	}
	if membershipType == "standard" {
		usuario.Type = "standard"
		usuario.Name = name
		usuario.MessageCharLimit = 100
		return
	}
	return
}
