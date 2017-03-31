package main

import (
	"./lib"
	"math/rand"
)

func main() {
	user_name_1, user_name_2 := lib.SetNames()

	user1 := lib.CreateUser(user_name_1, 100, rand.Intn(100))
	user2 := lib.CreateUser(user_name_2, 100, rand.Intn(100))

	lib.StartBattle(user1, user2);
}
