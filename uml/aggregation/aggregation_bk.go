package aggregation

//
//import (
//	"reflect"
//	"sync"
//)
//
//// University interface
//type University interface {
//	GetName() string
//	AddTeacher(teacher Human)
//}
//
//// MIT concrete
//type MIT struct {
//	name        string
//	teacherList []Human
//}
//
//func NewMIT(name string, teacherList []Human) University {
//	return &MIT{name: name, teacherList: teacherList}
//}
//
//func (this *MIT) GetName() string {
//	return this.name
//}
//
//func (this *MIT) AddTeacher(teacher Human) {
//	this.teacherList = append(this.teacherList, teacher)
//}
//
//// Human interface
//type Human interface {
//	GetName() string
//}
//
//// zhang3 concrete instance
//type zhang3 struct {
//}
//
//var zhang3_instance Human
//var zhang3_once sync.Once
//
//func NewZhang3() Human {
//	zhang3_once.Do(func() {
//		zhang3_instance = new(zhang3)
//	})
//	return zhang3_instance
//}
//
//func (this *zhang3) GetName() string {
//	return reflect.TypeOf((*zhang3)(nil)).Elem().Name()
//}
//
//// li4 concrete instance
//type li4 struct {
//}
//
//var li4_instance Human
//var li4_once sync.Once
//
//func NewLi4() Human {
//	li4_once.Do(func() {
//		li4_instance = new(li4)
//	})
//	return li4_instance
//}
//
//func (this *li4) GetName() string {
//	return reflect.TypeOf((*li4)(nil)).Elem().Name()
//}

//func Aggregation() {
//	mit := aggregation.NewMIT("MIT", nil)
//
//	zhang3 := aggregation.NewZhang3()
//	fmt.Println(zhang3.GetName())
//	li4 := aggregation.NewLi4()
//	fmt.Println(li4.GetName())
//
//	mit.AddTeacher(zhang3)
//	mit.AddTeacher(li4)
//
//	fmt.Println(mit.GetName())
//
//}
