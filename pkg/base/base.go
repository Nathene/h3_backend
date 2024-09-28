package base

type Base struct {
	Barracks *Barracks
	Camp     *Camp
	Smith    *Smith
}

func newBase() *Base {
	return &Base{
		Barracks: newBarracks(),
		Camp:     newCamp(),
		Smith:    newSmith(),
	}
}

type Barracks struct {
	Level  int
	Unlock map[string]bool
}

func newBarracks() *Barracks {
	return &Barracks{
		Level: 1,
		Unlock: map[string]bool{
			"Warrior": true,
			"Archer":  false,
		},
	}
}

type Camp struct {
	Level    int
	Capacity int
}

func newCamp() *Camp {
	return &Camp{
		Level:    1,
		Capacity: 10,
	}
}

type Smith struct {
	Level  int
	Troops map[string]int
}

func newSmith() *Smith {
	return &Smith{
		Level: 1,
		Troops: map[string]int{
			"Warrior": 1,
		},
	}
}

func CreateWarrior() *Troop {
	return &Troop{}
}

type Npc interface {
	Attack() (float64, error)
	Defend() (float64, error)
}
