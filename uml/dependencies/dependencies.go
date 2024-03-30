package dependencies

import "fmt"

// Human interface
type Human interface {
	Show()
}

// Teacher struct
type Teacher struct {
	name string
}

func NewTeacher(name string) Human {
	return &Teacher{name: name}
}

func (this *Teacher) Show() {
	fmt.Println(this.name)
}

// House interface
type House interface {
	Show(human Human)
}

// Classroom struct
type Classroom struct {
	name string
}

func NewClassroom(name string) House {
	return &Classroom{name: name}
}

func (this *Classroom) Show(human Human) {
	fmt.Println(this.name)
	human.Show()
}
