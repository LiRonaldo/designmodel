package template

import "fmt"

/**
模板模式 定义一个接口，一个方法按照一定的顺序去执行接口的方法，一个struct，实现了接口里的具体方法。
*/
type WorkInterface interface {
	GetUp()
	Work()
	Sleep()
}
type Worker struct {
	WorkInterface
}

func NewWorker(w WorkInterface) *Worker {
	return &Worker{w}
}
func (w *Worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Coder struct {
}

func (c *Coder) GetUp() {
	fmt.Println("coder is getup")
}
func (c *Coder) Work() {
	fmt.Println("coder is work")
}
func (c *Coder) Sleep() {
	fmt.Println("coder is sleep")
}
