package strategy

import "fmt"

type WeaponStrategy interface {
	UseWeapon() //使用武器
}

type AK47 struct {
}

func (ak *AK47) UseWeapon() {
	fmt.Println("Use AK47")
}

type Knife struct {
}

func (kn *Knife) UseWeapon() {
	fmt.Println("Use Knife")
}
