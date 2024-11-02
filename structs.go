package main

import (
	"example.com/structs/user"
)

func main() {
	firstName := user.GetUserData("Please enter your first name: ")
	lastName := user.GetUserData("Please enter your last name: ")
	birthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		return
	}

	admin := user.NewAdmin("dummy@gmail.com", "dummy-password")

	admin.User.OutputUserDetails()

	// ... do something awesome with that gathered data!
	appUser.OutputUserDetails()

	appUser.ClearUserName()
	appUser.OutputUserDetails()
}
