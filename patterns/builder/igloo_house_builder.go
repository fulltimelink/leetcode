package builder

type iglooHouseBuilder struct {
	House
}

func newIglooHouseBuilder() *iglooHouseBuilder {
	return &iglooHouseBuilder{}
}

func (b *iglooHouseBuilder) setWindowType() {
	b.WindowType = "Snow Window"
}

func (b *iglooHouseBuilder) setDoorType() {
	b.DoorType = "Snow Door"
}

func (b *iglooHouseBuilder) setNumFloor() {
	b.Floor = 1
}

func (b *iglooHouseBuilder) getHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
