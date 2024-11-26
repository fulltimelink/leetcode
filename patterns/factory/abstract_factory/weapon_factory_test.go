package abstract_factory

import "testing"

func TestGetWeaponFactory(t *testing.T) {
	rf := GetWeaponFactory("red")
	if rf == nil {
		t.Error("nil factory")
	}
	rf.MakeShortGun("ak47").Shoot()
}
