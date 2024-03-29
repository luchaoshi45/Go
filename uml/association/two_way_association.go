package association

import "fmt"

// Human 抽象接口
type Human interface {
	HaveFun()
	GetName() string
	SetName(Name string)
	SetLovers(Lovers Human)
}

// Girl 实现
type Girl struct {
	Name   string
	Lovers Human
}

func NewGirl(Name string, Lovers Human) Human {
	return &Girl{Name: Name, Lovers: Lovers}
}

func (g *Girl) HaveFun() {
	fmt.Println(g.GetName(), " Have Fun With ", g.Lovers.GetName())
}

func (g *Girl) GetName() string {
	return g.Name
}

func (g *Girl) SetName(Name string) {
	g.Name = Name
}

func (g *Girl) SetLovers(Lovers Human) {
	g.Lovers = Lovers
}

// Boy 实现
type Boy struct {
	Name   string
	Lovers Human
}

func NewBoy(Name string, Lovers Human) Human {
	return &Boy{Name: Name, Lovers: Lovers}
}

func (b *Boy) HaveFun() {
	fmt.Println(b.GetName(), " Have Fun With ", b.Lovers.GetName())
}

func (b *Boy) GetName() string {
	return b.Name
}

func (b *Boy) SetName(Name string) {
	b.Name = Name
}

func (b *Boy) SetLovers(Lovers Human) {
	b.Lovers = Lovers
}
