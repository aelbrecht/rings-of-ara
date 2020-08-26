package textures

import (
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"
)

type Texture struct {
	image *image.RGBA
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
		image: t.image.SubImage(image.Rect(x*16, y*16, (x+1)*16, (y+1)*16)).(*image.RGBA),
	}
}
