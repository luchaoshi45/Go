package test

import (
	"Go/uml/association"
)

func TwoWayAssociation() {
	girl := association.NewGirl("girl1", nil)
	boy := association.NewBoy("boy1", nil)
	girl.SetLovers(boy)
	boy.SetLovers(girl)

	girl.HaveFun()
	boy.HaveFun()

}
