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

	w.Player.Vel.Y += 0.6
	if w.Player.Vel.Y > 6 {
		w.Player.Vel.Y = 6
	}
	if w.Player.Pos.Y >= 0 {
		w.Player.Pos.Y = 0
		w.Player.Vel.Y = 0
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
			if w.Player.Pos.Y == 0 {
				w.Player.Vel.Y = -10
			}
		}
	}

	w.Player.Pos.X += int64(math.Round(w.Player.Vel.X))
	w.Player.Pos.Y += int64(math.Round(w.Player.Vel.Y))

}
