package events

import (
	"math"
	"rings-of-ara/internal/world"
)

func HandleEvents(w *world.Model, container *EventQueue) {

	w.Player.Step += 1

	if w.Player.Vel.X > 0.2 {
		w.Player.Vel.X += -0.2
	} else if w.Player.Vel.X < -0.2 {
		w.Player.Vel.X += 0.2
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

	for container.Size > 0 {
		container.Size--
		event := container.Items[container.Size]

		switch event.Kind {
		case Left:
			w.Player.Vel.X += -0.75
			if w.Player.Vel.X < -4 {
				w.Player.Vel.X = -4
			}
			w.Player.Direction = -1
			break
		case Right:
			w.Player.Vel.X += 0.75
			if w.Player.Vel.X > 4 {
				w.Player.Vel.X = 4
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

	if MaskCollision(newAbsX, w.Player.Pos.Y, pw, ph, w) {
		w.Player.Vel.X /= 2
		if math.Abs(w.Player.Vel.X) < 0.5 {
			w.Player.Vel.X = 0
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
		return b.Kind != 0
	}
	return false
}
