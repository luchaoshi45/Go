package test

import "Go/uml/association"

func Association() {
	book := association.NewBook("cv")
	reader := association.NewReader("lcs", book)
	reader.ReaderBook()
}
