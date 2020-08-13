package proxy

import "testing"

func TestTaskProxy_RentHouse(t *testing.T) {
	task := NewTaskProxy()
	task.RentHouse("盖房子花了", 10000)
}
