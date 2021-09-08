package dronetown

type Location interface {
	GetID() string
	GetRoad() chan Drone
}

type Address struct {
	ID   string
	Road chan Drone
}

func (d Address) GetID() string {
	return d.ID
}

func (d Address) GetRoad() chan Drone {
	return d.Road
}
