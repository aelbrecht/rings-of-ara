package characters

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"rings-of-ara/internal/textures"
	"rings-of-ara/internal/world"
)

func PlayerSprite(ch *world.Character, screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	sprite := textures.SpriteAlycia.Image()
	if ch.Direction < 0 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(16, 0)
	}

	op.GeoM.Scale(3, 3)

	op.GeoM.Translate(300, 200)

	i := 0
	j := (ch.Step / 10) % 6

	if ch.Vel.X != 0 {
		i = 1
		j = (ch.Step / 7) % 6
	}

	if ch.Pos.Y < 0 {
		i = 2
		j = 3
		if ch.Vel.Y > 0 {
			if ch.Pos.Y > -15 {
				j = 1
			} else if ch.Pos.Y > -30 {
				j = 2
			}
		} else {
			if ch.Pos.Y > -15 {
				j = 5
			} else if ch.Pos.Y > -30 {
				j = 4
			}
		}
	}

	xf, yf := ch.Pos.ToFloat64()
	op.GeoM.Translate(xf, yf)

	// create an ebiten image as textures are imported as rgba go images
	img2, _ := ebiten.NewImageFromImage(sprite.SubImage(image.Rect(16*j, i*28, 16*(j+1), (i+1)*28)), ebiten.FilterDefault)
	_ = screen.DrawImage(img2, op)
}
