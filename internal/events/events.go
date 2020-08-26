package events

import "rings-of-ara/internal/world"

func HandleEvents(w *world.Model, container *EventQueue) {

	w.Player.Step += 1

	w.Player.Vel.X = 0

	w.Player.Vel.Y += 0.1
	if w.Player.Vel.Y > 2 {
		w.Player.Vel.Y = 2
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
			w.Player.Vel.X = -1.5
			w.Player.Direction = -1
			break
		case Right:
			w.Player.Vel.X = 1.5
			w.Player.Direction = 1
			break
		case Jump:
			if w.Player.Pos.Y == 0 {
				w.Player.Vel.Y = -3
			}
		}
	}

	w.Player.Pos.X += int64(w.Player.Vel.X)
	w.Player.Pos.Y += int64(w.Player.Vel.Y)

}
