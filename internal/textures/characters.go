package textures

var (
	CharacterHairs *TileSet
	CharacterHands *TileSet
	CharacterFeet  *TileSet
	CharacterHeads *TileSet
	CharacterLimbs *TileSet
	CharacterUpper *TileSet
	CharacterLower *TileSet
)

func init() {
	dir := AssetsDir + "/sprites/character"
	CharacterHairs = LoadTileSet(dir+"/hair.png", 32, 32)
	CharacterHands = LoadTileSet(dir+"/hand.png", 6, 6)
	CharacterFeet = LoadTileSet(dir+"/feet.png", 6, 6)
	CharacterHeads = LoadTileSet(dir+"/head.png", 18, 18)
	CharacterLimbs = LoadTileSet(dir+"/limb.png", 4, 10)
	CharacterUpper = LoadTileSet(dir+"/upper.png", 32, 24)
	CharacterLower = LoadTileSet(dir+"/lower.png", 14, 6)
}
