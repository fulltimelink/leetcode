package abstract_factory

import "fmt"

type RedAk47 struct {
	ShortGun
}

func (g *RedAk47) Shoot() {
	fmt.Println("red ak47 short gun shoot")
}
