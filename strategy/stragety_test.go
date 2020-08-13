package strategy

import "testing"

func TestContext_SetStragety(t *testing.T) {
	context := NewContext()
	strageA := NewStragetyA()
	context.SetStragety(strageA)
	context.Execute()
	StragetyB := NewStragetyB()
	context.SetStragety(StragetyB)
	context.Execute()
}
