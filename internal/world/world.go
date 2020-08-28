package world

type Block struct {
	Kind uint16
}

type Chunk struct {
	Data [ChunkSize * ChunkSize]Block
}

type Planet struct {
	Size   uint32
	Chunks map[ChunkPosition]*Chunk
}

type Model struct {
	Camera *Camera
	Player *Character
	Planet *Planet
}
