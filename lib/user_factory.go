package lib

import "fmt"

func CreateUser() *user {
	userName:= SetUserName()

	return &user{name: userName, health: 100, protection: 100}
}

func SetUserName() string {
	var userName string

	fmt.Print("Set user name:")
	fmt.Scanf("%s", &userName)

	return userName
}