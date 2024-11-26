package abstract_factory

import "log"

type BlueFactory struct{}

func (b *BlueFactory) MakeShortGun(weapon string) IShortGun {
	if "ak47" == weapon {
		return &BlueAk47{
			ShortGun{
				name:  "Blue AK47 gun",
				power: 4,
			},
		}
	}
	if "musket" == weapon {
		return &BlueMusket{
			ShortGun{
				name:  "Blue Musket gun",
				power: 32,
			},
		}
	}
	log.Fatalf("BlueFactory can not make %s short gun!\n", weapon)
	return nil
}
