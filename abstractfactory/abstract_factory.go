package abstractfactory

import "fmt"

/**
抽象工厂模式
*/
type Lunch interface {
	Cook()
}
type Rise struct {
}

func (r *Rise) Cook() {
	fmt.Println("it is rise.")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("it is tomato.")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {
}

func (s SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}
