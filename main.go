package main

import (
	"fmt"
	"math/rand"
	"./lib"
)

func main() {
	user_name_1, user_name_2 := lib.SetNames()

	user1 := lib.CreteUser(user_name_1, 100, rand.Intn(100))
	user2 := lib.CreteUser(user_name_2, 100, rand.Intn(100))

	stopBattle := true

	for stopBattle {
		step(user1)
		step(user2)
		if user1.GetHealth() <= 1 || user1.GetHealth() <= 1 {
			stopBattle = false
		}
	}
}

func step(u lib.Action) int {
	BitPower := rand.Intn(u.GetHealth())
	u.Biter(BitPower)

	fmt.Println(u.GetName(), ": bit_power->", BitPower, " health->", u.GetHealth())

	return BitPower
}
