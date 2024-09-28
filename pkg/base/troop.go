package base

const (
	Warrior = iota // 0
	Archer         // 1
	Mage           // 2
)

type Troop struct {
	Name  string
	Stats Stats
	Class Class
}

type Class struct {
	Name     string
	Skills   []skill
	Passive  []passive
	Space    int
	SubClass subClass
}

type subClass struct {
	Name    string
	Passive []passive
}

type skill struct {
	Name   string
	Action action
}

type passive struct {
	Name  string
	Stats Stats
}

type action struct {
	Name     string
	Damage   int
	Accuracy int
}
