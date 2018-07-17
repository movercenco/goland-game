package main

import (
	"./lib"
)

func main() {
	user1 := lib.CreateUser()
	user2 := lib.CreateUser()

	lib.StartBattle(user1, user2)
}
