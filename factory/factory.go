package factory

import "fmt"

/**
工厂模式
*/
type Restaurant interface {
	GetFood()
}
type DongLaiShun struct {
}

func (d *DongLaiShun) GetFood() {
	fmt.Println("东来顺的饭菜已经生产完毕，继续…………；")
}

type QianFeng struct {
}

func (q *QianFeng) GetFood() {
	fmt.Println("庆丰包子饭菜已经生产完毕,继续…………；")
}
func NewRestaurant(name string) Restaurant {
	switch name {
	case "Q":
		return &QianFeng{}
	case "D":
		return &DongLaiShun{}
	}
	return nil
}
