package textures

import "github.com/hajimehoshi/ebiten"

var TileSetGrassland *ebiten.Image
var TileSetUnderground *ebiten.Image
var TileSetWorldInterface *ebiten.Image
var TileSetWeapons *ebiten.Image

const AssetsDir = "./assets"

func init() {
	TileSetGrassland = LoadImage(AssetsDir + "/world/grassland.png")
	TileSetUnderground = LoadImage(AssetsDir + "/world/underground.png")
	TileSetWorldInterface = LoadImage(AssetsDir + "/interface/world.png")
	TileSetWeapons = LoadImage(AssetsDir + "/sprites/weapons.png")
}