package lib

type user struct {
	name       string
	health     int
	protection int
}

type UserAction interface {
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
	if u.protection >= power{
		u.protection -= power
	} else {
		if u.protection != 0 {
			u.health = u.health - (power - u.protection)
			u.protection = 0
		} else{
			u.health = u.health - power
		}
	}
}