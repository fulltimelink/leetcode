package abstract_factory

type IWeaponFactory interface {
	MakeShortGun(string) IShortGun
}

func GetWeaponFactory(f string) IWeaponFactory {
	if "red" == f {
		return &RedFactory{}
	}
	if "blue" == f {
		return &BlueFactory{}
	}
	return nil
}
