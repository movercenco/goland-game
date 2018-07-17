package lib

import (
	"fmt"
	"math/rand"
)

func StartBattle(user1 UserAction, user2 UserAction) {
	round:= 1

	for {
		fmt.Println("###### Round ", round, " ######")

		Step(user1)
		if user1.GetHealth() <= 1 {
			ShowWinner(user2)
			break
		}

		Step(user2)
		if user2.GetHealth() <= 1 {
			ShowWinner(user2)
			break
		}
		fmt.Println()
		round++
	}
}

func Step(u UserAction) {
	bitPower := rand.Intn(100)
	u.Biter(bitPower)
	fmt.Println(u.GetName(), ": bit power ", bitPower, ", health ", u.GetHealth(), ", protection ", u.GetProtection())
}

func ShowWinner(u UserAction)  {
	fmt.Println("The winner is ", u.GetName())
}