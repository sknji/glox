package value

import (
	"fmt"
	"github.com/urijn/glox/shared"
)

type Value struct {
	ValType ValueType
	Val     *ConcreteValue
}

func NewValue(valueType ValueType, val interface{}) *Value {
	return &Value{ValType: valueType, Val: NewConcreteValue(val)}
}

func (v *Value) Is(t ValueType) bool {
	return v.ValType == t
}

func (v *Value) String() string {
	if v == nil {
		return ""
	}

	return fmt.Sprintf("%+v", v.Val)
}

func (v *Value) Print() {
	fmt.Print(v.String())
}

func (v *Value) Println() {
	fmt.Println(v.String())
}

type ValueStore struct {
	Values []*Value
	Count  int
	Cap    int
}

func NewValueStore() *ValueStore {
	return &ValueStore{Values: make([]*Value, shared.DefaultCapacity)}
}

func (c *ValueStore) Write(value *Value) {
	if c.Cap < c.Count+1 {
		c.Cap = shared.GrowCapacity(c.Cap)
		tmp := make([]*Value, c.Cap)
		c.Values = tmp
	}

	c.Values[c.Count] = value
	c.Count += 1
}

func (c *ValueStore) Free() {
	c.Values = nil
	c.Cap = 0
	c.Count = 0
}
