package builder

type IHouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func GetHouseBuilder(htype string) IHouseBuilder {
	if htype == "normal" {
		return newNormalHouseBuilder()
	}
	if htype == "igloo" {
		return newIglooHouseBuilder()
	}
	return nil
}
