package singleton

import (
	"fmt"
	"sync"
)

/* 懒汉式
Singleton（单例）：在单例类的内部实现只生成一个实例，
同时它提供一个静态的getInstance()工厂方法，让客户可以访问它的唯一实例；
为了防止在外部对其实例化，将其构造函数设计为私有；
在单例类内部定义了一个Singleton类型的静态对象，作为外部共享的唯一实例。
单例模式要解决的问题是：
保证一个类永远只能有一个对象，且该对象的功能依然能被其他模块使用。
*/

/*  ===== v1 =====
// 首字母小写，类外不能调用
type singleton struct {
	name string
}

// 指针指向唯一对象，指针永远不能改变值
var instance *singleton = new(singleton)

// 返回 instance 的普通公有函数
func GetInstance() *singleton {
	return instance
}

func (st *singleton) Show() {
	fmt.Println("singleton Show")
}
*/

/* ===== v2 =====
type singleton struct {
	name string // 空结构体的指针 new 多少次都相同
}

var instance *singleton
var lock sync.Mutex
var initialized uint32

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (st *singleton) Show() {
	fmt.Println("singleton Show")
}
*/

/* ===== v3 ===== */
type singleton struct {
	name string // 空结构体的指针 new 多少次都相同
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
func (st *singleton) Show() {
	fmt.Println("singleton Show")
}
