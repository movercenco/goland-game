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
	GetProtection() int
}

func (u *user) GetHealth() int {
	return u.health
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetProtection() int {
	return u.protection
}

func (u *user) Biter(power int) {
	//pow := power
	if u.protection > 0 {
		u.protection -= power;
	}

	if u.protection == 0 {
		u.health = u.health - power;
	}

	if u.protection < 0 {
		u.health = u.health + u.protection;
		u.protection = 0;
	}
}

//CreateUser return new instance of User
func CreateUser(name string, health int, protection int) *user {
	return &user{name: name, health: health, protection: protection}
}
