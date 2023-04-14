package main

func Go(vehicle actions) {
	if sled, ok := vehicle.(*Sled); ok {
		sled.pushStart()
	} else {
		vehicle.startEngine()
	}

	vehicle.drive()
}

type actions interface {
	drive()
	startEngine()
}

type Vehicle struct{}

func (v Vehicle) drive() {
	// TODO: implement
}

func (v Vehicle) startEngine() {
	// TODO: implement
}

func (v Vehicle) pushStart() {
	// TODO: implement
}

type Car struct {
	Vehicle
}

type Sled struct {
	Vehicle
}

func (s Sled) startEngine() {
	//
}

func (s Sled) stopEngine() {
	//
}

func (s Sled) pushStart() {
	//
}
