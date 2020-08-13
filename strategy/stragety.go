package strategy

/**
策略模式
*/
import "fmt"

type Stragety interface {
	Execute()
}
type StragetyA struct {
}
type StragetyB struct {
}

func (s *StragetyA) Execute() {
	fmt.Println("StragetyA executed")
}
func NewStragetyA() Stragety {
	return &StragetyA{}
}
func (s *StragetyB) Execute() {
	fmt.Println("StragetyB executed")
}
func NewStragetyB() Stragety {
	return &StragetyB{}
}

type Context struct {
	stragety Stragety
}

func NewContext() *Context {
	return &Context{}
}
func (c *Context) SetStragety(stragety Stragety) {
	c.stragety = stragety
}
func (c *Context) Execute() {
	c.stragety.Execute()
}
