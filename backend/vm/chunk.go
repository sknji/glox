package vm

type Chunk struct {
	Code  []byte
	Lines []uint
	Count int
	Cap   int

	Constants *ValueStore
}

func NewChunk() *Chunk {
	return &Chunk{
		Count:     0,
		Cap:       0,
		Code:      []byte{},
		Lines:     []uint{},
		Constants: NewValueStore(),
	}
}

func (c *Chunk) Write(byte byte, line uint) {
	if c.Cap < c.Count+1 {
		c.Cap = GrowCapacity(c.Cap)

		tmp := make([]uint8, c.Cap)
		copy(tmp, c.Code)
		c.Code = tmp

		tmp1 := make([]uint, c.Cap)
		copy(tmp1, c.Lines)
		c.Lines = tmp1
	}

	c.Code[c.Count] = byte
	c.Lines[c.Count] = line
	c.Count += 1
}

func (c *Chunk) Free() {
	c.Code = nil
	c.Lines = nil
	c.Cap = 0
	c.Count = 0
}

func (c *Chunk) AddConstant(value Value) uint8 {
	c.Constants.Write(value)
	return uint8(c.Constants.Count - 1)
}
