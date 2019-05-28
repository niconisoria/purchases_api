package models

import "fmt"

type User struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	DNI      int    `json:"dni"`
	ID       int64  `json:"user_id"`
}

func (u *User) FullName() string {
	return fmt.Sprintf("%v %v", u.Name, u.LastName)
}

//The reason why the endPoint didn't allow to add new purchases.

/*func (u *User) IsValid() bool {
	return false
}*/

func (u *User) IsValid() bool {
	return u.DNI >= 1111111
}
