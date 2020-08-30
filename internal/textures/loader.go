package textures

var SpriteAlycia *Texture
var GrasslandTileMap *Texture
var UndergroundTileMap *Texture

const AssetsDir = "./assets/"

func init() {
	SpriteAlycia = LoadTexture(AssetsDir + "sprites/alycia.png")
	GrasslandTileMap = LoadTexture(AssetsDir + "world/grassland.png")
	UndergroundTileMap = LoadTexture(AssetsDir + "world/underground.png")
}
