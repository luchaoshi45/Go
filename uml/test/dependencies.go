package test

import "Go/uml/dependencies"

func Dependencies() {
	teacher := dependencies.NewTeacher("zhnag3")
	cl301 := dependencies.NewClassroom("301")
	cl301.Show(teacher)
}
