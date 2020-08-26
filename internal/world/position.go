package world

type Coordinates struct {
	X, Y int64
}

func (c Coordinates) ToFloat64() (float64, float64) {
	return float64(c.X), float64(c.Y)
}

type Vector struct {
	X, Y float64
}
