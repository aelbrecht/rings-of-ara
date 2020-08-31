package textures

import "github.com/hajimehoshi/ebiten"

type TileSet = *ebiten.Image

var SpriteAlycia *Texture
var GrasslandTileMap *Texture
var UndergroundTileMap *Texture
var TileSetGrassland TileSet
var TileSetUnderground TileSet

const AssetsDir = "./assets/"

func init() {
	SpriteAlycia = LoadTexture(AssetsDir + "sprites/alycia.png")
	GrasslandTileMap = LoadTexture(AssetsDir + "world/grassland.png")
	UndergroundTileMap = LoadTexture(AssetsDir + "world/underground.png")
	TileSetGrassland = LoadTileSet(AssetsDir + "world/grassland.png")
	TileSetUnderground = LoadTileSet(AssetsDir + "world/underground.png")
}
