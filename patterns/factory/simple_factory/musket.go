package simple_factory

import "fmt"

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun{
			name:  "Musket",
			power: 1,
		},
	}
}

func (musket *Musket) shoot() {
	fmt.Println("musket gun shoot")
}
