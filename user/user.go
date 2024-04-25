package user

import "github.com/srswiggy/first-api/role"

type User struct {
	Name  string
	Email string
	Role  role.Role
}

func CreateUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
		Role:  role.User,
	}
}

func CreateAmdin(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
		Role:  role.Admin,
	}
}

func CreateSpecialUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
		Role:  role.SpecialUser,
	}
}
