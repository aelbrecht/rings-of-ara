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
	tex, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Printf("could not load texture %s\n", src)
		log.Fatal(err)
	}
	return &Texture{
		image: tex,
	}
}

func (t *Texture) Image() *ebiten.Image {
	return t.image
}
