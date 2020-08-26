package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
	"rings-of-ara/internal/draw/characters"
	"rings-of-ara/internal/events"
	"rings-of-ara/internal/world"
)

type GameScreen struct {
	H int
	W int
}

type GameProperties struct {
	Screen GameScreen
}

// main game object
type Game struct {
	Props          GameProperties
	InputHandler   func(*world.Model, *events.EventQueue)
	EventHandler   func(*world.Model, *events.EventQueue)
	EventContainer *events.EventQueue
	World          *world.Model
}

// game update step
func (g *Game) Update(_ *ebiten.Image) error {

	// handle input
	g.InputHandler(g.World, g.EventContainer)

	// handle events
	g.EventHandler(g.World, g.EventContainer)

	return nil
}

// prevents screen resizing
func (g *Game) Layout(_, _ int) (int, int) {
	return g.Props.Screen.W, g.Props.Screen.H
}

// draw loop
// uses a buffer to make drawing cleaner from different routines
func (g *Game) Draw(screen *ebiten.Image) {
	_ = screen.Fill(color.RGBA{228, 241, 254, 255})
	d := g.World.Player.Draw
	d(g.World.Player, screen)
}

func main() {

	player := &world.Character{
		Pos:  world.Coordinates{},
		Vel:  world.Vector{},
		Draw: characters.PlayerSprite,
	}

	g := &Game{
		EventContainer: events.MakeEventContainer(),
		Props: GameProperties{
			Screen: GameScreen{
				H: 400,
				W: 600,
			},
		},
		EventHandler: events.HandleEvents,
		InputHandler: events.HandleGameInput,
		World: &world.Model{
			Player: player,
		},
	}

	// set parameters and start loop
	ebiten.SetWindowSize(g.Props.Screen.W*2, g.Props.Screen.H*2)
	ebiten.SetWindowTitle("Rings of Ara")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
