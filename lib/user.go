package lib

type user struct {
	name       string
	health     int
	protection int
}

// Action user
type Action interface {
	Biter(int)
	GetHealth() int
	GetName() string
}

func (u *user) Biter(power int) {
	//pow := power
	if u.protection != 0 {

	}

	u.health = u.health - power
}

func (u *user) GetHealth() int {
	return u.health
}

func (u *user) GetName() string {
	return u.name
}

//CreteUser retrun new instance of User
func CreteUser(name string, health int, protection int) *user {
	user := &user{name: name, health: health, protection: protection}
	return user
}
