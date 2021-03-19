package value

import (
	"fmt"
	"github.com/urijn/glox/shared"
)


type Value float64

func (v Value) String() string {
	return fmt.Sprintf("%g", v)
}

func (v Value) Print() {
	fmt.Print(v.String())
}

func (v Value) PrintLn() {
	fmt.Println(v.String())
}

type ValueStore struct {
	Values []Value
	Count  int
	Cap    int
}

func NewValueStore() *ValueStore {
	return &ValueStore{Values: make([]Value, shared.DefaultCapacity)}
}

func (c *ValueStore) Write(value Value) {
	if c.Cap < c.Count+1 {
		c.Cap = shared.GrowCapacity(c.Cap)
		tmp := make([]Value, c.Cap)
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
