package events

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"math"
	"rings-of-ara/internal/textures"
	"rings-of-ara/internal/world"
)

var activeBlock uint16 = 0

func HandleEvents(w *world.Model, container *EventQueue) {

	w.Player.Step += 1

	if w.Player.Vel.X > 0.19 {
		w.Player.Vel.X -= 0.12
	} else if w.Player.Vel.X < -0.19 {
		w.Player.Vel.X += 0.12
	} else {
		w.Player.Vel.X = 0
	}

	pw, ph := w.Player.Mask.PixelValues()

	onSolid := MaskOnSolid(w.Player.Pos.X, w.Player.Pos.Y-1, pw, ph, w)
	if !onSolid {
		w.Player.Vel.Y -= 0.6
		if w.Player.Vel.Y < -15 {
			w.Player.Vel.Y = -15
		}
	}

	maxV := 4.0

	for container.Size > 0 {
		container.Size--
		event := container.Items[container.Size]

		switch event.Kind {
		case Move:
			data := event.Data.([]int)
			target := w.Camera.ToWorld(world.Coordinates{int64(data[0]), int64(data[1])})
			w.Player.Target = target
			break
		case Key0:
			activeBlock = 0
			break
		case Key1:
			activeBlock = 1
			break
		case Primary:
			// add/remove blocks

			// get block pos from cursor
			pos := event.Data.([]int)
			curWorldPos := w.Camera.ToWorld(world.Coordinates{int64(pos[0]), int64(pos[1])})

			// set block
			chunk := w.Planet.GetChunk(curWorldPos)
			var t *ebiten.Image
			if activeBlock == 1 {
				t = textures.TileSetUnderground.SubImage(image.Rectangle{
					Min: image.Point{0, 0},
					Max: image.Point{10, 10},
				}).(*ebiten.Image)
			}
			// TODO: use generator to set blocks so that we don't have to manually set texture data
			chunk.Data[curWorldPos.ToRelativeBlockPosition().Index()] = world.Block{
				Solid:   activeBlock != 0,
				Kind:    activeBlock,
				TexMain: t,
			}
			break
		case Shift:
			maxV = 6
			break
		case Alt:
			w.Player.Vel.Y = 0
			break
		case Left:
			w.Player.Vel.X += -0.2
			if w.Player.Vel.X < -maxV {
				w.Player.Vel.X = -maxV
			}
			w.Player.Direction = -1
			break
		case Right:
			w.Player.Vel.X += 0.2
			if w.Player.Vel.X > maxV {
				w.Player.Vel.X = maxV
			}
			w.Player.Direction = 1
			break
		case Jump:
			if onSolid {
				w.Player.Vel.Y = 11
			}
			break
		case Up:
			w.Player.Vel.Y = 4
			break
		case Down:
			w.Player.Vel.Y = -4
			break
		}
	}

	if w.Player.Vel.X == 0 && w.Player.Vel.Y == 0 {
		return
	}

	newAbsX := w.Player.Pos.X + int64(math.Round(w.Player.Vel.X))
	newAbsY := w.Player.Pos.Y + int64(math.Round(w.Player.Vel.Y))

	jumped := false
	if w.Player.Vel.X < 0 {
		b1 := w.Planet.GetBlock(world.Coordinates{w.Player.Pos.X - world.BlockPixelSize, w.Player.Pos.Y - int64(w.Player.Mask.H)*3 + world.BlockPixelSize + 1})
		b2 := w.Planet.GetBlock(world.Coordinates{w.Player.Pos.X - world.BlockPixelSize, w.Player.Pos.Y - int64(w.Player.Mask.H)*3 + world.BlockPixelSize/2})
		if onSolid && b1 != nil && b2 != nil && !b1.Solid && b2.Solid {
			w.Player.Vel.Y = 7
			jumped = true
		}
	} else if w.Player.Vel.X > 0 {
		b1 := w.Planet.GetBlock(world.Coordinates{w.Player.Pos.X + int64(pw) + world.BlockPixelSize, w.Player.Pos.Y - int64(w.Player.Mask.H)*3 + world.BlockPixelSize + 1})
		b2 := w.Planet.GetBlock(world.Coordinates{w.Player.Pos.X + int64(pw) + world.BlockPixelSize, w.Player.Pos.Y - int64(w.Player.Mask.H)*3 + world.BlockPixelSize/2})
		if onSolid && b1 != nil && b2 != nil && !b1.Solid && b2.Solid {
			w.Player.Vel.Y = 7
			jumped = true
		}
	}

	if !MaskCollision(newAbsX, newAbsY, pw, ph, w) {
		w.Player.Pos.X = newAbsX
		w.Player.Pos.Y = newAbsY
		return
	}

	if MaskCollision(newAbsX, w.Player.Pos.Y, pw, ph, w) {
		if !jumped {
			w.Player.Vel.X /= 2
			if math.Abs(w.Player.Vel.X) < 0.5 {
				w.Player.Vel.X = 0
			}
		}
	} else {
		w.Player.Pos.X = newAbsX
	}

	if MaskCollision(w.Player.Pos.X, newAbsY, pw, ph, w) {
		if w.Player.Vel.Y < 0 {
			for i := 0; i < 10; i++ {
				if MaskOnSolid(w.Player.Pos.X, w.Player.Pos.Y-int64(i)-1, pw, ph, w) {
					w.Player.Pos.Y -= int64(i)
					w.Player.Vel.Y = 0
					return
				}
			}
		}
		w.Player.Vel.Y /= 2
		if math.Abs(w.Player.Vel.Y) < 0.5 {
			w.Player.Vel.Y = 0
		}
	} else {
		w.Player.Pos.Y = newAbsY
	}
}

func MaskOnSolid(x int64, y int64, pw int, ph int, w *world.Model) bool {
	collided := false
	for dx := 0; dx < pw && !collided; dx += world.BlockPixelSize {
		dc := world.Coordinates{x + int64(dx), y - int64(ph)}
		if CheckCollision(dc, w) {
			collided = true
		}
	}
	dc := world.Coordinates{x + int64(pw), y - int64(ph)}
	if CheckCollision(dc, w) {
		collided = true
	}
	return collided
}

func MaskCollision(x int64, y int64, pw int, ph int, w *world.Model) bool {
	collided := false
	for dx := 0; dx < pw && !collided; dx += world.BlockPixelSize {
		for dy := 0; dy < ph && !collided; dy += world.BlockPixelSize {
			dc := world.Coordinates{x + int64(dx), y - int64(dy)}
			if CheckCollision(dc, w) {
				collided = true
			}
		}
	}
	for dy := 0; dy < ph && !collided; dy += world.BlockPixelSize {
		dc := world.Coordinates{x + int64(pw), y - int64(dy)}
		if CheckCollision(dc, w) {
			collided = true
		}
	}
	for dx := 0; dx < pw && !collided; dx += world.BlockPixelSize {
		dc := world.Coordinates{x + int64(dx), y - int64(ph)}
		if CheckCollision(dc, w) {
			collided = true
		}
	}
	dc := world.Coordinates{x + int64(pw), y - int64(ph)}
	if CheckCollision(dc, w) {
		collided = true
	}
	return collided
}

func CheckCollision(p world.Coordinates, w *world.Model) bool {
	b := w.Planet.GetBlock(p)
	if b != nil {
		return b.Solid
	}
	return false
}
