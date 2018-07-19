package main

import "combat/lib"

func main() {
	user1 := lib.CreateUser()
	user2 := lib.CreateUser()

	lib.StartBattle(user1, user2)
}
