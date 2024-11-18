package simple_factory

import "errors"

// --  @# 抽离共有方法为接口
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
	shoot()
}

// --  @# 工厂创建方法，要啥创建啥
func GetGun(gunType string) (IGun, error) {
	if "ak47" == gunType {
		return newAk47(), nil
	}
	if "musket" == gunType {
		return newMusket(), nil
	}
	return nil, errors.New("wrong gun type passed")
}
