package main

func Go(vehicle actions) {
	vehicle.start()
	vehicle.drive()
}

type actions interface {
	start()
	drive()
}

type Cart struct {
	poweredActions
}

func (c Cart) start() {
	c.poweredActions.startEngine()
}

func (c Car) drive() {
	// TODO: implement
}

type poweredActions interface{}

type Buggy struct{}

func (b Buggy) start() {
	// 시작
}

func (b Buggy) drive() {
	// TODO: implement
}
