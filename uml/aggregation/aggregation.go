package aggregation

// University interface
type University interface {
	GetName() string
	AddTeacher(teacher Human)
}

// MIT concrete
type MIT struct {
	name        string
	teacherList []Human
}

func NewMIT(name string, teacherList []Human) University {
	return &MIT{name: name, teacherList: teacherList}
}

func (this *MIT) GetName() string {
	return this.name
}

func (this *MIT) AddTeacher(teacher Human) {
	this.teacherList = append(this.teacherList, teacher)
}

// Human interface
type Human interface {
	GetName() string
}

// Teacher concrete
type Teacher struct {
	name string
}

func NewTeacher(name string) Human {
	return &Teacher{name: name}
}

func (this *Teacher) GetName() string {
	return this.name
}
