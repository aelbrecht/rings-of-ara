package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	"rings-of-ara/internal/draw"
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

var tmpimg, _ = ebiten.NewImage(1280, 800, ebiten.FilterDefault)

// draw loop
// uses a buffer to make drawing cleaner from different routines
func (g *Game) Draw(screen *ebiten.Image) {

	// fill sky
	_ = screen.Fill(color.RGBA{228, 241, 254, 255})

	// fill block layer
	draw.Planet(g.World, screen)

	d := g.World.Player.Draw
	d(g.World.Player, g.World, screen)

	chunkDebug := ""
	activeChunks := g.World.Camera.VisibleChunks()
	for _, chunk := range activeChunks {
		chunkDebug += fmt.Sprintf("%d,%d ", chunk.X, chunk.Y)
	}

	_ = ebitenutil.DebugPrint(screen,
		fmt.Sprintf("TPS: %0.2f\nPosition: %f,%f\n%s",
			ebiten.CurrentTPS(),
			float64(g.World.Player.Pos.X),
			float64(g.World.Player.Pos.Y),
			chunkDebug,
		),
	)
}

func main() {

	player := &world.Character{
		Pos:  world.Coordinates{10000, 0},
		Vel:  world.Vector{},
		Draw: draw.PlayerSprite,
	}

	g := &Game{
		EventContainer: events.MakeEventContainer(),
		Props: GameProperties{
			Screen: GameScreen{
				H: 800,
				W: 1280,
			},
		},
		EventHandler: events.HandleEvents,
		InputHandler: events.HandleGameInput,
		World: &world.Model{
			Camera: &world.Camera{
				Subject: player,
				Size:    world.Rectangle{W: 1280, H: 800},
			},
			Player: player,
			Planet: &world.Planet{
				Size:   5000,
				Chunks: make(map[world.ChunkPosition]*world.Chunk),
			},
		},
	}

	// set parameters and start loop
	ebiten.SetWindowSize(g.Props.Screen.W, g.Props.Screen.H)
	ebiten.SetWindowTitle("Rings of Ara")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
