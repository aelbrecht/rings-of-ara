package world

type Coordinates struct {
	X, Y int64
}

func (c Coordinates) Add(b Coordinates) Coordinates {
	return Coordinates{
		X: c.X + b.X,
		Y: c.Y + b.Y,
	}
}

func (c Coordinates) Sub(b Coordinates) Coordinates {
	return Coordinates{
		X: c.X - b.X,
		Y: c.Y - b.Y,
	}
}

type Rectangle struct {
	W, H int64
}

type ChunkPosition struct {
	X, Y uint32
}

type BlockPosition struct {
	X, Y uint32
}

func (c Coordinates) ToChunkPosition() ChunkPosition {
	return c.ToBlockPosition().ToChunkPosition()
}

func (c BlockPosition) ToChunkPosition() ChunkPosition {
	return ChunkPosition{
		X: c.X / 32,
		Y: c.Y / 32,
	}
}

func (c Coordinates) ToBlockPosition() BlockPosition {
	return BlockPosition{
		X: uint32(c.X / (32 * 3)),
		Y: uint32(c.Y / (32 * 3)),
	}
}

func (c Coordinates) ToFloat64() (float64, float64) {
	return float64(c.X), float64(c.Y)
}

type Vector struct {
	X, Y float64
}
