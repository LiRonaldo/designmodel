package abstractfactory

import "testing"

func TestNewSimpleLunchFactory(t *testing.T) {
	LunchFactory := NewSimpleLunchFactory()
	rise := LunchFactory.CreateFood()
	rise.Cook()
	veg := LunchFactory.CreateVegetable()
	veg.Cook()
}
