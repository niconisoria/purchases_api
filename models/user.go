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

func (u *User) IsValid() bool {
	return u.DNI >= 1111111
}
