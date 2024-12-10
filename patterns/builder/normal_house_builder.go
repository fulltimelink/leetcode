package builder

type normalHouseBuilder struct {
	House
}

func newNormalHouseBuilder() *normalHouseBuilder {
	return &normalHouseBuilder{}
}

func (b *normalHouseBuilder) setWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *normalHouseBuilder) setDoorType() {
	b.DoorType = "Wooden Door"
}

func (b *normalHouseBuilder) setNumFloor() {
	b.Floor = 2
}

func (b *normalHouseBuilder) getHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
