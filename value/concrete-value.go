package value

import (
	"fmt"
)

type ConcreteValue struct {
	Val interface{}
}

func NewConcreteValue(val interface{}) *ConcreteValue {
	return &ConcreteValue{Val: val}
}

func (c *ConcreteValue) Get() interface{} {
	return c.Val
}

func (c *ConcreteValue) GetObject() Object {
	return (c.Val).(Object)
}

func (c *ConcreteValue) GetBool() bool {
	return (c.Val).(bool)
}

func (c *ConcreteValue) GetNumber() float64 {
	return (c.Val).(float64)
}

func (c *ConcreteValue) Print() {
	fmt.Printf(c.String())
}

func (c *ConcreteValue) String() string {
	return fmt.Sprintf("%+v", c.Val)
}
