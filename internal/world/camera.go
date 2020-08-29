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
	w2 := int64(ChunkPixelSize)
	h2 := int64(ChunkPixelSize)

	chunks = addChunkToList(Coordinates{-w2, -h2}.Add(c.Subject.Pos).ToChunkPosition(), chunks) // top left
	chunks = addChunkToList(Coordinates{0, -h2}.Add(c.Subject.Pos).ToChunkPosition(), chunks)   // top middle
	chunks = addChunkToList(Coordinates{w2, -h2}.Add(c.Subject.Pos).ToChunkPosition(), chunks)  // top right

	chunks = addChunkToList(Coordinates{-w2, h2}.Add(c.Subject.Pos).ToChunkPosition(), chunks) // bottom left
	chunks = addChunkToList(Coordinates{0, h2}.Add(c.Subject.Pos).ToChunkPosition(), chunks)   // bottom middle
	chunks = addChunkToList(Coordinates{w2, h2}.Add(c.Subject.Pos).ToChunkPosition(), chunks)  // bottom right

	chunks = addChunkToList(Coordinates{-w2, 0}.Add(c.Subject.Pos).ToChunkPosition(), chunks) // center top
	chunks = addChunkToList(Coordinates{0, 0}.Add(c.Subject.Pos).ToChunkPosition(), chunks)   // center middle
	chunks = addChunkToList(Coordinates{w2, 0}.Add(c.Subject.Pos).ToChunkPosition(), chunks)  // center bottom

	return chunks
}

func (c *Camera) ToWorld(coords Coordinates) Coordinates {
	coords.Y = -coords.Y
	newPos := c.Subject.Pos
	newPos = newPos.Sub(Coordinates{
		X: c.Size.W/2,
		Y: -c.Size.H/2,
	})
	newPos = coords.Add(newPos)
	return newPos
}

func (c *Camera) ToScreen(coords Coordinates) Coordinates {
	newPos := c.Subject.Pos
	newPos = newPos.Sub(Coordinates{
		X: c.Size.W/2,
		Y: -c.Size.H/2,
	})
	newPos = coords.Sub(newPos)
	newPos.Y = -newPos.Y
	return newPos
}
