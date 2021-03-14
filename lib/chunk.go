package lib

type Chunk struct {
	code  []byte
	count int
	cap   int

	consts *ValueStore
}

func NewChunk() *Chunk {
	return &Chunk{
		count:  0,
		cap:    0,
		code:   make([]byte, DefaultCapacity),
		consts: NewValueStore(),
	}
}

func (c *Chunk) Write(byte byte) {
	if c.cap < c.count+1 {
		c.cap = cap(c.code)
		c.code = c.code[:c.cap]
	}

	c.code[c.count] = byte
	c.count += 1
}

func (c *Chunk) Free() {
	c.code = nil
	c.cap = 0
	c.count = 0
}

func (c *Chunk) AddConstant(value Value) uint8 {
	c.consts.Write(value)
	return uint8(c.consts.count - 1)
}
