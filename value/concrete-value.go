package value

import "fmt"

type ConcreteValue struct {
	Val interface{}
}

func NewConcreteValue(val interface{}) *ConcreteValue {
	return &ConcreteValue{Val: val}
}

func (c *ConcreteValue) Get() interface{} {
	return c.Val
}

func (c *ConcreteValue) Set(val interface{}) {
	c.Val = val
}

func (c *ConcreteValue) GetAsBool() bool {
	return (c.Val).(bool)
}

func (c *ConcreteValue) GetAsNumber() float64 {
	return (c.Val).(float64)
}

func (c *ConcreteValue) Print() {
	fmt.Printf("%+v", c.Val)
}

func (c *ConcreteValue) String() string {
	return fmt.Sprintf("%+v", c.Val)
}
