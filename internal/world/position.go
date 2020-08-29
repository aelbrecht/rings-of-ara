package world

const (
	BlockPixelSize   = 16 * 3
	ChunkSize        = 16
	ChunkPixelSize   = BlockPixelSize * ChunkSize
	SeaLevel         = ChunkSize * 100
	UnderGroundLevel = ChunkSize * 80
	SkyLevel         = ChunkSize * 120
)

type Coordinates struct {
	X, Y int64
}

func (c Coordinates) Add(b Coordinates) Coordinates {
	return Coordinates{
		X: c.X + b.X,
		Y: c.Y + b.Y,
	}
}

func (c Coordinates) Inv() Coordinates {
	return Coordinates{
		X: -c.X,
		Y: -c.Y,
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

type CharacterMask struct {
	W, H int
}

func (m CharacterMask) PixelValues() (int, int) {
	return m.W * 3, m.H * 3
}

type ChunkPosition struct {
	X, Y uint32
}

type BlockPosition struct {
	X, Y uint32
}

type RelativeBlockPosition struct {
	X, Y int
}

func (pos RelativeBlockPosition) Values() (int, int) {
	return pos.X, pos.Y
}

func BlockIndexToPosition(i int) RelativeBlockPosition {
	y := i / ChunkSize
	x := i - y*ChunkSize
	return RelativeBlockPosition{
		X: x,
		Y: y,
	}
}

func BlockPositionToIndex(p RelativeBlockPosition) int {
	return p.Y*ChunkSize + p.X
}

func (c Coordinates) ToChunkPosition() ChunkPosition {
	return c.ToBlockPosition().ToChunkPosition()
}

func (c Coordinates) ToRelativeBlockPosition() RelativeBlockPosition {
	p := c.ToBlockPosition()
	rp := RelativeBlockPosition{
		X: int(p.X % ChunkSize),
		Y: int(p.Y % ChunkSize),
	}
	return rp
}

func (c BlockPosition) ToChunkPosition() ChunkPosition {
	return ChunkPosition{
		X: c.X / ChunkSize,
		Y: c.Y / ChunkSize,
	}
}

func (c ChunkPosition) ToCoordinates() Coordinates {
	return Coordinates{
		X: int64(c.X) * ChunkPixelSize,
		Y: int64(c.Y) * ChunkPixelSize,
	}
}

func (c Coordinates) ToBlockPosition() BlockPosition {
	return BlockPosition{
		X: uint32(c.X / BlockPixelSize),
		Y: uint32(c.Y / BlockPixelSize),
	}
}

func (c Coordinates) ValuesFloat() (float64, float64) {
	return float64(c.X), float64(c.Y)
}

func (c Coordinates) Values() (int64, int64) {
	return c.X, c.Y
}

type Vector struct {
	X, Y float64
}
