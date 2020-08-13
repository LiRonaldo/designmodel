package template

import "testing"

func TestWorker_Daily(t *testing.T) {
	f := NewWorker(&Coder{})
	f.Daily()
}
