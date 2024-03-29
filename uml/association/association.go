package association

import "fmt"

// AbstractReader 抽象读者 1
type AbstractReader interface {
	ReaderBook()
}

// AbstractBook 抽象书 2
type AbstractBook interface {
	GetName() string
	SetName(Name string)
}

// Book 具体书 3
type Book struct {
	Name string
}

func NewBook(Name string) AbstractBook {
	return &Book{Name: Name}
}

func (b *Book) GetName() string {
	return b.Name
}

func (b *Book) SetName(Name string) {
	b.Name = Name
}

// Reader 具体读者 4
// 一般关联关系  Book 不构成 Reader的一部分
type Reader struct {
	Name string
	Book AbstractBook
}

func NewReader(Name string, Book AbstractBook) AbstractReader {
	return &Reader{Name: Name, Book: Book}
}

func (r *Reader) ReaderBook() {
	fmt.Println(r.Name, " Reader ", r.Book.GetName())
}
