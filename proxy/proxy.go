package proxy

import "fmt"

/**
代理模式
*/
type ITask interface {
	RentHouse(desc string, price int)
}
type Task struct {
}

func (t *Task) RentHouse(desc string, price int) {
	fmt.Println(desc, price)
}

type TaskProxy struct {
	task *Task
}

func NewTaskProxy() *TaskProxy {
	return &TaskProxy{task: &Task{}}
}
func (t *TaskProxy) RentHouse(desc string, price int) {
	t.task.RentHouse(desc, price)
}
