package abstract_factory

import "fmt"

type BlueMusket struct {
	ShortGun
}

func (g *BlueMusket) Shoot() {
	fmt.Println("blue musket short gun shoot")
}
