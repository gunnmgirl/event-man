package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Jo",
			lastName:  "Jo",
			birthDate: "10/10/2002",
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	(*u).firstName = ""
	(*u).lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("please fill in the required values")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
