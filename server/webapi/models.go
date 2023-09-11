package webapi

import "fmt"

type Account struct {
	Id       string
	Email    string
	password string
}

func (a Account) String() string {
	return fmt.Sprintf("id=%v,  Email=%v",
		a.Id, a.Email)
}
