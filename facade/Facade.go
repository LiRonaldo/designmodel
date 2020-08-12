package facade

import "fmt"

/**
门面模式  ,有一个对外的门面，外边的人，只需要提需求，不需要知道门面里边是怎么做的。
*/
type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}
func (c *CarModel) SetModel() {
	fmt.Println("CarModel……")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}
func (c *CarEngine) setEngine() {
	fmt.Println("CarEngine……")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}
func (c *CarBody) setCarBody() {
	fmt.Println("CarBody……")
}

type CarFacde struct {
	model  *CarModel
	engine *CarEngine
	body   *CarBody
}

func NewCarFacde() *CarFacde {
	return &CarFacde{model: NewCarModel(), engine: NewCarEngine(), body: NewCarBody()}
}
func (c *CarFacde) GreateCar() {
	c.model.SetModel()
	c.engine.setEngine()
	c.body.setCarBody()
}
