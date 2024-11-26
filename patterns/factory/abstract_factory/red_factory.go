package abstract_factory

import "log"

type RedFactory struct {
}

func (s *RedFactory) MakeShortGun(weapon string) IShortGun {
	if "ak47" == weapon {
		return &BlueAk47{
			ShortGun{
				name:  "Red AK47 gun",
				power: 4,
			},
		}
	}
	if "musket" == weapon {
		return &RedMusket{
			ShortGun{
				name:  "Red Musket gun",
				power: 32,
			},
		}
	}
	log.Fatalf("RedFactory can not make %s short gun!\n", weapon)
	return nil
}
