package simple_factory

import "testing"

func TestGetGun(t *testing.T) {
	ak47, _ := GetGun("ak47")
	musket, _ := GetGun("musket")
	ak47.shoot()
	musket.shoot()
}
