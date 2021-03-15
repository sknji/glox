package stack

type Chunk struct {
	code  []byte
	lines []uint
	count int
	cap   int

	constants *ValueStore
}

func NewChunk() *Chunk {
	return &Chunk{
		count:     0,
		cap:       0,
		code:      make([]byte, DefaultCapacity),
		lines:     make([]uint, DefaultCapacity),
		constants: NewValueStore(),
	}
}

func (c *Chunk) Write(byte byte, line uint) {
	if c.cap < c.count+1 {
		c.cap = GrowCapacity(c.cap)

		tmp := make([]uint8, c.cap)
		copy(tmp, c.code)
		c.code = tmp

		tmp1 := make([]uint, c.cap)
		copy(tmp1, c.lines)
		c.lines = tmp1
	}

	c.code[c.count] = byte
	c.lines[c.count] = line
	c.count += 1
}

func (c *Chunk) Free() {
	c.code = nil
	c.lines = nil
	c.cap = 0
	c.count = 0
}

func (c *Chunk) AddConstant(value Value) uint8 {
	c.constants.Write(value)
	return uint8(c.constants.count - 1)
}
