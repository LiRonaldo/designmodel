package facade

import "testing"

func TestCarFacde_GreateCar(t *testing.T) {
	creatCar := NewCarFacde()
	creatCar.GreateCar()
}
