package bad

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
	// 아무런 동작도 하지 않도록 오버라이드를 수행한다.
}

func (s Sled) stopEngine() {
	// 아무런 동작도 하지 않도록 오버라이드를 수행한다.
}

func (s Sled) pushStart() {
	// TODO: implement
}
