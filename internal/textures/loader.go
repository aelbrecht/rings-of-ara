package textures

import "github.com/hajimehoshi/ebiten"

type TileSet = *ebiten.Image

var SpriteAlycia *Texture
var TileSetGrassland TileSet
var TileSetUnderground TileSet
var TileSetWorldInterface TileSet
var TileSetWeapons TileSet

const AssetsDir = "./assets/"

func init() {
	SpriteAlycia = LoadTexture(AssetsDir + "sprites/alycia.png")
	TileSetGrassland = LoadTileSet(AssetsDir + "world/grassland.png")
	TileSetUnderground = LoadTileSet(AssetsDir + "world/underground.png")
	TileSetWorldInterface = LoadTileSet(AssetsDir + "interface/world.png")
	TileSetWeapons = LoadTileSet(AssetsDir + "sprites/weapons.png")
}
