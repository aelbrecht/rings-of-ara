package draw

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"rings-of-ara/internal/world"
)

func aimLine(ch *world.Character, cam *world.Camera, dst *ebiten.Image) {

	ox, oy := cam.ToScreen(ch.Pos).ValuesFloat()
	tx, ty := cam.ToScreen(ch.Target).ValuesFloat()

	ebitenutil.DrawLine(dst, ox, oy, tx, ty, color.RGBA{0, 0, 0, 0})
}