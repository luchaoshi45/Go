package test

import "Go/designPatterns/behavioral/strategy"

func Strategy() {
	hero := strategy.Hero{}

	hero.SetWeaponStrategy(new(strategy.AK47))
	hero.Fight()

	hero.SetWeaponStrategy(new(strategy.Knife))
	hero.Fight()
}
