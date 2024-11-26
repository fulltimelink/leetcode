package abstract_factory

import "fmt"

type RedMusket struct {
	ShortGun
}

func (g *RedMusket) Shoot() {
	fmt.Println("red musket short gun shoot")
}
