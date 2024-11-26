package abstract_factory

import "fmt"

type BlueAk47 struct {
	ShortGun
}

func (g *BlueAk47) Shoot() {
	fmt.Println("blue ak47 short gun shoot")
}
