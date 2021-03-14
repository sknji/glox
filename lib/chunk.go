package lib

type Chunk struct {
	count    int
	capacity int

	code []uint8
}

func NewChunk() *Chunk {
	return &Chunk{count: 0, capacity: 0, code: []uint8{}}
}

func (c *Chunk) Write(byte uint8) {
	if c.capacity < c.count + 1 {
		c.capacity = GrowCapacity(c.capacity)
		c.code = c.code[:c.capacity]
	}

	c.code[c.count] = byte
	c.count += 1
}

func (c *Chunk) free() {
	c.code = nil
	c.capacity = 0
	c.count = 0
}
