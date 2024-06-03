package decorate

import "fmt"

type FilmDecorate struct {
	ap AbstractPhone
}

func NewFilmDecorate(ap AbstractPhone) AbstractDecorate {
	return &FilmDecorate{ap: ap}
}

func (fd *FilmDecorate) Show() {
	fd.ap.Show()
	fmt.Println("FilmDecorate")
}
