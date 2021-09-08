package dronetown

type Shop struct {
	ID       string
	Provides func(string) bool
	Road     chan Drone
	Prov     string
}

func (s Shop) GetID() string {
	return s.ID
}

func (s Shop) GetRoad() chan Drone {
	return s.Road
}
