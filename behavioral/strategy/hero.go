package strategy

type Hero struct {
	strategy WeaponStrategy
}

func (h *Hero) SetWeaponStrategy(ws WeaponStrategy) {
	h.strategy = ws
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}
