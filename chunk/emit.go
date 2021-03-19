package chunk

func (c *Chunk) EmitBytes(line uint, bytes ...byte) {
	for _, b := range bytes {
		c.Write(b, line)
	}
}
