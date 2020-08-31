package textures

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"
	"rings-of-ara/internal/world"
)

type Texture struct {
	image *image.RGBA
}

func LoadTileSet(src string) *ebiten.Image {
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

// loads a texture from disk into an rgba image
func LoadTexture(src string) *Texture {

	// load image from disk
	raw, err := os.Open(src)
	img, _, err := image.Decode(raw)
	if err != nil {
		log.Printf("failed to load %s\n", src)
		log.Fatal(err)
	}

	// convert image to rgba texture
	b := img.Bounds()
	tex := image.NewRGBA(img.Bounds())
	draw.Draw(tex, b, img, b.Min, draw.Src)

	return &Texture{
		image: tex,
	}
}

func (t *Texture) Image() *image.RGBA {
	return t.image
}

func (t *Texture) Tile(x int, y int) Texture {
	return Texture{
		image: t.image.SubImage(image.Rect(x*world.BlockSize, y*world.BlockSize, (x+1)*world.BlockSize, (y+1)*world.BlockSize)).(*image.RGBA),
	}
}
