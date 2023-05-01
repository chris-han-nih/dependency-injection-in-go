package main

func Go(vehicle actions) {
	switch concrete := vehicle.(type) {
	case poweredActions:
		concrete.startEngine()
	case unpoweredActions:
		concrete.pushStart()
	}

	vehicle.drive()
}

type actions interface {
	drive()
}

type poweredActions interface {
	actions
	startEngine()
	stopEngine()
}

type unpoweredActions interface {
	actions
	pushStart()
}

type Vehicle struct{}

func (v Vehicle) drive() {
	// TODO: implement
}

type PoweredVehicle struct {
	Vehicle
}

func (v PoweredVehicle) startEngine() {
	// 일반적인 엔진 시작 코드
}

type Car struct {
	PoweredVehicle
}

type Buggy struct {
	Vehicle
}

func (b Buggy) pushStart() {
	// 아무런 동작도 하지 않는다.
}
