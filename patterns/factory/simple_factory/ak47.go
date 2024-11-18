package simple_factory

import "fmt"

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

func (g *Ak47) shoot() {
	fmt.Println("ak47 gun shoot")
}
