package abstract_factory

import "fmt"

// --  @# 抽离共有方法为接口
type IShortGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
	Shoot()
}

type ShortGun struct {
	name  string
	power int
}

func (g *ShortGun) Shoot() {
	fmt.Println("gun shoot")
}

func (g *ShortGun) setName(name string) {
	g.name = name
}

func (g *ShortGun) setPower(power int) {
	g.power = power
}

func (g *ShortGun) getName() string {
	return g.name
}

func (g *ShortGun) getPower() int {
	return g.power
}
