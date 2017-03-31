package lib

import (
	"fmt"
	"math/rand"
)

//SetNamesFromConsole retrun new instance of User
func SetNames() (string, string) {
	var (
		user_name_1 string
		user_name_2 string
	);

	fmt.Print("First user:")
	fmt.Scanf("%s", &user_name_1)
	fmt.Print("Second user:")
	fmt.Scanf("%s", &user_name_2)

	return user_name_1, user_name_2
}

// Step return bit power for one step
func Step(u Action) {
	BitPower := rand.Intn(u.GetHealth())
	u.Biter(BitPower)

	fmt.Println(u.GetName(), ": bit_power->", BitPower, " health->", u.GetHealth(), " GetProtection->", u.GetProtection())
}

//StartBattle
func StartBattle(user1 Action, user2 Action) {

	for {
		Step(user1)
		if user1.GetHealth() <= 1 {
			break
		}

		Step(user2)
		if user2.GetHealth() <= 1 {
			break
		}
	}
}
