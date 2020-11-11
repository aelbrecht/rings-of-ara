package textures

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"log"
	"os"
)

type Texture struct {
	image *ebiten.Image
}

type TileMap struct {
	ts *TileSet
	x  int
	y  int
}

func (tm *TileMap) GetTexture() *ebiten.Image {
	return tm.ts.GetTile(tm.x, tm.y)
}

func MakeTileMap(x int, y int, tileSet *TileSet) TileMap {
	return TileMap{
		ts: tileSet,
		x:  x,
		y:  y,
	}
}

type TileSet struct {
	image *ebiten.Image
	w     int
	h     int
}

func (t *TileSet) GetTileMap(x int, y int) TileMap {
	return MakeTileMap(x, y, t)
}

func (t *TileSet) GetTile(x int, y int) *ebiten.Image {
	sub := t.image.SubImage(image.Rectangle{
		Min: image.Point{X: x * t.w, Y: y * t.h},
		Max: image.Point{X: (x + 1) * t.w, Y: (y + 1) * t.h},
	})
	return sub.(*ebiten.Image)
}

func LoadImage(src string) *ebiten.Image {
	raw, err := os.Open(src)
	img, _, err := image.Decode(raw)
	if err != nil {
		log.Printf("failed to load %s\n", src)
		log.Fatal(err)
	}
	tex, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Printf("could not load texture %s\n", src)
		log.Fatal(err)
	}
	return tex
}

func LoadTileSet(src string, w int, h int) *TileSet {
	img := LoadImage(src)
	return &TileSet{
		image: img,
		w:     w,
		h:     h,
	}
}

// loads a texture from disk into an rgba image
func LoadTexture(src string) *Texture {
	tex := LoadImage(src)
	return &Texture{
		image: tex,
	}
}

func (t *Texture) Image() *ebiten.Image {
	return t.image
}
