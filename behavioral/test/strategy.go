package test

import "Go/behavioral/strategy"

func Strategy() {
	hero := strategy.Hero{}

	hero.SetWeaponStrategy(new(strategy.AK47))
	hero.Fight()

	hero.SetWeaponStrategy(new(strategy.Knife))
	hero.Fight()
}
