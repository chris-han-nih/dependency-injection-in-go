package lsp

func Go(vehicle actions) {
	vehicle.start()
	vehicle.drive()
}

type actions interface {
	start()
	drive()
}

type Car struct {
	poweredVehicle
}

func (c Car) start() {
	c.poweredVehicle.startEngine()
}

func (c Car) drive() {
	// TODO: implement
}

type poweredVehicle struct{}

func (p poweredVehicle) startEngine() {
	// 일반적인 엔진 시작 코드
}

type Buggy struct{}

func (b Buggy) start() {
	// 시작
}

func (b Buggy) drive() {
	// TODO: implement
}
