package lib

import "fmt"

type Value float64

func (v Value) String() string {
	return fmt.Sprintf("%g", v)
}

type ValueStore struct {
	values []Value
	count  int
	cap    int
}

func NewValueStore() *ValueStore {
	return &ValueStore{values: make([]Value, DefaultCapacity)}
}

func (c *ValueStore) Write(value Value) {
	if c.cap < c.count+1 {
		c.cap = cap(c.values)
		c.values = c.values[:c.cap]
	}

	c.values[c.count] = value
	c.count += 1
}

func (c *ValueStore) Free() {
	c.values = nil
	c.cap = 0
	c.count = 0
}
