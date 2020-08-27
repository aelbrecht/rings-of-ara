package world

type Camera struct {
	Subject *Character
	Size    Rectangle
}

func addChunkToList(c ChunkPosition, l []ChunkPosition) []ChunkPosition {
	for _, p := range l {
		if c == p {
			return l
		}
	}
	l = append(l, c)
	return l
}

func (c *Camera) VisibleChunks() []ChunkPosition {
	chunks := make([]ChunkPosition, 0)
	w2 := c.Size.W / 2
	h2 := c.Size.H / 2

	chunks = addChunkToList(Coordinates{-w2, -h2}.Add(c.Offset()).ToChunkPosition(), chunks) // top left
	chunks = addChunkToList(Coordinates{-w2, h2}.Add(c.Offset()).ToChunkPosition(), chunks)  // bottom left
	chunks = addChunkToList(Coordinates{0, -h2}.Add(c.Offset()).ToChunkPosition(), chunks)   // middle top
	chunks = addChunkToList(Coordinates{0, h2}.Add(c.Offset()).ToChunkPosition(), chunks)    // middle bottom
	chunks = addChunkToList(Coordinates{w2, -h2}.Add(c.Offset()).ToChunkPosition(), chunks)  // top right
	chunks = addChunkToList(Coordinates{w2, h2}.Add(c.Offset()).ToChunkPosition(), chunks)   // bottom right

	return chunks
}

func (c *Camera) Offset() Coordinates {
	return c.Subject.Pos.Add(Coordinates{
		X: -c.Size.W/2 + (16 * 3 / 2),
		Y: -c.Size.H/2 + (28 * 3 / 2),
	})
}
