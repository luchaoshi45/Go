package test

import "Go/uml/inheritance"

func Inheritance() {
	humanFac := inheritance.HumanFac{}
	human := humanFac.CreateAnimal("human")
	human.Eat()

	teacherFac := inheritance.TeacherFac{}
	teacher := teacherFac.CreateAnimal("human")
	teacher.Eat()

	studentFac := inheritance.StudentFac{}
	student := studentFac.CreateAnimal("human")
	student.Eat()
}
