package events

import (
	"github.com/hajimehoshi/ebiten"
	"rings-of-ara/internal/world"
)

type state struct {
	K          [10]bool
	Inventory  bool
	Use        bool
	Reload     bool
	Enter      bool
	Map        bool
	Shift      bool
	Jump       bool
	Primary    bool
	Secondary  bool
	LastCursor []int
}

var keyMap map[int]string

func init() {
	keyMap = make(map[int]string)
	keyMap[0] = Key0
	keyMap[1] = Key1
	keyMap[2] = Key2
	keyMap[3] = Key3
	keyMap[4] = Key4
	keyMap[5] = Key5
	keyMap[6] = Key6
	keyMap[7] = Key7
	keyMap[8] = Key8
	keyMap[9] = Key9
}

var s = state{
	K:          [10]bool{false, false, false, false, false, false, false, false, false, false},
	Inventory:  false,
	Use:        false,
	Reload:     false,
	Enter:      false,
	Map:        false,
	Shift:      false,
	Jump:       false,
	Primary:    false,
	Secondary:  false,
	LastCursor: []int{0, 0},
}

func HandleGenericMouseInput(es *EventQueue) {
	_, dy := ebiten.Wheel()
	if dy != 0 {
		es.Add(Wheel, dy)
	}
	x, y := ebiten.CursorPosition()
	if x != s.LastCursor[0] || y != s.LastCursor[1] {
		es.Add(Move, []int{x, y})
		s.LastCursor = []int{x, y}
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		es.Add(PrimaryDown, []int{x, y})
		s.Primary = true
	} else if s.Primary {
		es.Add(Primary, []int{x, y})
		s.Primary = false
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		es.Add(SecondaryDown, []int{x, y})
		s.Secondary = true
	} else if s.Secondary {
		es.Add(Secondary, []int{x, y})
		s.Secondary = false
	}
}

func HandleGameInput(w *world.Model, es *EventQueue) {

	// num keys
	j := 0
	for i := ebiten.Key0; i <= ebiten.Key9; i++ {
		if ebiten.IsKeyPressed(i) {
			s.K[j] = true
		} else if s.K[j] {
			es.Add(keyMap[j], nil)
			s.K[j] = false
		}
		j++
	}

	// handle movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		es.Add(Up, nil)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		es.Add(Down, nil)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		es.Add(Left, nil)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		es.Add(Right, nil)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		es.Add(Fly, nil)
	}

	// handle actions
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		s.Inventory = true
	} else if s.Inventory {
		es.Add(Inventory, nil)
		s.Inventory = false
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !s.Jump {
		es.Add(Jump, nil)
		s.Jump = true
	} else if s.Jump && !ebiten.IsKeyPressed(ebiten.KeySpace) {
		s.Jump = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		s.Use = true
	} else if s.Use {
		es.Add(Use, nil)
		s.Use = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		s.Reload = true
	} else if s.Reload {
		es.Add(Reload, nil)
		s.Reload = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		s.Enter = true
	} else if s.Enter {
		es.Add(Enter, nil)
		s.Enter = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyM) {
		s.Map = true
	} else if s.Map {
		es.Add(Map, nil)
		s.Map = false
	}

	// handle modifiers
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		es.Add(Shift, nil)
	}
	if ebiten.IsKeyPressed(ebiten.KeyAlt) {
		es.Add(Alt, nil)
	}

	// handle mouse
	HandleGenericMouseInput(es)

}
