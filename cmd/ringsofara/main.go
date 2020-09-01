package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"rings-of-ara/internal/blocks"
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

type GameBuffers struct {
	BlockBackLayer  *ebiten.Image
	BlockFrontLayer *ebiten.Image
}

// main game object
type Game struct {
	Props          GameProperties
	InputHandler   func(*world.Model, *events.EventQueue)
	EventHandler   func(*world.Model, *events.EventQueue)
	EventContainer *events.EventQueue
	World          *world.Model
	Buffers        GameBuffers
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

func main() {

	screen := GameScreen{
		H: 800,
		W: 1280,
	}

	player := &world.Character{
		Mask: world.CharacterMask{16, 28},
		Pos:  world.Coordinates{world.ChunkPixelSize * 1000, world.ChunkPixelSize * 105},
		Vel:  world.Vector{},
		Draw: draw.PlayerSprite,
	}

	bufferBlockBackLayer, _ := ebiten.NewImage(screen.W, screen.H, ebiten.FilterDefault)
	bufferBlockFrontLayer, _ := ebiten.NewImage(screen.W, screen.H, ebiten.FilterDefault)

	g := &Game{
		EventContainer: events.MakeEventContainer(),
		Props: GameProperties{
			Screen: screen,
		},
		EventHandler: events.HandleEvents,
		InputHandler: events.HandleGameInput,
		World: &world.Model{
			Camera: &world.Camera{
				Subject: player,
				Size:    world.Rectangle{W: int64(screen.W), H: int64(screen.H)},
			},
			Player: player,
			Planet: &world.Planet{
				Size:   5000,
				Chunks: make(map[world.ChunkPosition]*world.Chunk),
			},
		},
		Buffers: GameBuffers{
			BlockBackLayer:  bufferBlockBackLayer,
			BlockFrontLayer: bufferBlockFrontLayer,
		},
	}

	go blocks.ChunkGenerator(g.World)

	// set parameters and start loop
	ebiten.SetWindowSize(g.Props.Screen.W, g.Props.Screen.H)
	ebiten.SetWindowTitle("Rings of Ara")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
