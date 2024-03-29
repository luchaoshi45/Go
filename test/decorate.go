package test

import "Go/decorate"

func Decorate() {
	hwp := decorate.NewHuaWeiPhone("GB200")
	hwp.Show()

	xiaomi := decorate.NewXiaoMiPhone("X200")
	xiaomi.Show()

	xiaomi = decorate.NewFilmDecorate(xiaomi)
	xiaomi.Show()
}
