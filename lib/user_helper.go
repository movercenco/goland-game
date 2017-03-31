package lib

import "fmt"

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
