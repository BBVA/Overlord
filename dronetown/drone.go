package dronetown

type Stamp struct {
	ID   string
	Item string
}

type Drone struct {
	ID       string
	Passport []Stamp
	Backpack []string
	Brain    Stack
}
