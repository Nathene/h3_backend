package base

type Player struct {
	Name   string
	Troops []Troop
	Base   Base
	Level  int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Base:  *newBase(),
		Level: 1,
	}
}
