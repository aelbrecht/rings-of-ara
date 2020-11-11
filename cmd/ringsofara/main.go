package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"rings-of-ara/internal/blocks"
	"rings-of-ara/internal/events"
	"rings-of-ara/internal/textures"
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

	alysia := world.CharacterFeature{
		Texture:  textures.CharacterUpper.GetTile(0, 0),
		Rotation: 0,
		OffsetX:  64,
		OffsetY:  64,
		AnchorX:  16,
		AnchorY:  12,
		ChildrenBefore: []world.CharacterFeature{
			{
				Texture:        textures.CharacterHeads.GetTile(0, 0),
				Rotation:       0,
				OffsetX:        1,
				OffsetY:        -6,
				AnchorX:        9,
				AnchorY:        18 - 3,
				ChildrenBefore: nil,
			}, {
				Texture:  textures.CharacterUpper.GetTile(1, 0),
				Rotation: 0,
				OffsetX:  0,
				OffsetY:  0,
				AnchorX:  16,
				AnchorY:  12,
			}, {
				Texture:  textures.CharacterUpper.GetTile(2, 0),
				Rotation: 0,
				OffsetX:  0,
				OffsetY:  0,
				AnchorX:  16,
				AnchorY:  12,
			}, {
				Texture:        textures.CharacterHairs.GetTile(0, 0),
				Rotation:       0,
				OffsetX:        1,
				OffsetY:        -6,
				AnchorX:        16,
				AnchorY:        32 - 7 - 3,
				ChildrenBefore: nil,
			},{
				Texture:        nil,
				Rotation:       0,
				OffsetX:        0,
				OffsetY:        11,
				AnchorX:        7,
				AnchorY:        1,
				ChildrenBefore: []world.CharacterFeature{
					{
						Texture:  textures.CharacterLimbs.GetTile(1, 1),
						Rotation: 0,
						OffsetX:  -3,
						OffsetY:  2,
						AnchorX:  1,
						AnchorY:  1,
						ChildrenBehind: []world.CharacterFeature{
							{
								Texture:        textures.CharacterLimbs.GetTile(1, 1),
								Rotation:       0,
								OffsetX:        0,
								OffsetY:        7,
								AnchorX:        1,
								AnchorY:        1,
								ChildrenBefore: []world.CharacterFeature{
									{
										Texture:        textures.CharacterFeet.GetTile(1, 0),
										Rotation:       0,
										OffsetX:        0,
										OffsetY:        8,
										AnchorX:        1,
										AnchorY:        3,
										ChildrenBefore: nil,
									},
								},
							},
						},
					},
				},
			}, {
				Texture:  textures.CharacterLimbs.GetTile(2, 0),
				Rotation: 0.4,
				OffsetX:  -8,
				OffsetY:  -2,
				AnchorX:  1,
				AnchorY:  1,
				ChildrenBehind: []world.CharacterFeature{
					{
						Texture:        textures.CharacterLimbs.GetTile(0, 0),
						Rotation:       -1,
						OffsetX:        0,
						OffsetY:        7,
						AnchorX:        1,
						AnchorY:        1,
						ChildrenBefore: nil,
					},
				},
			},
		},
		ChildrenBehind: []world.CharacterFeature{{
			Texture:  textures.CharacterLimbs.GetTile(2, 0),
			Rotation: -0.4,
			OffsetX:  5,
			OffsetY:  -2,
			AnchorX:  1,
			AnchorY:  1,
			Darken:   0.02,
			ChildrenBehind: []world.CharacterFeature{
				{
					Texture:        textures.CharacterLimbs.GetTile(0, 0),
					Rotation:       1,
					OffsetX:        0,
					OffsetY:        7,
					AnchorX:        1,
					AnchorY:        1,
					ChildrenBefore: nil,
					Darken:         0.02,
				},
			},
		},
			{
				Texture:        textures.CharacterHairs.GetTile(0, 1),
				Rotation:       0,
				OffsetX:        1,
				OffsetY:        -6,
				AnchorX:        16,
				AnchorY:        32 - 7 - 3,
				ChildrenBefore: nil,
			},
			{
				Texture:  textures.CharacterLower.GetTile(1, 0),
				Rotation: 0,
				OffsetX:  0,
				OffsetY:  11,
				AnchorX:  7,
				AnchorY:  1,
				ChildrenBehind: []world.CharacterFeature{
					{
						Texture:  textures.CharacterLimbs.GetTile(1, 1),
						Rotation: 0,
						OffsetX:  2,
						OffsetY:  2,
						AnchorX:  1,
						AnchorY:  1,
						Darken:   0.02,
						ChildrenBehind: []world.CharacterFeature{
							{
								Texture:        textures.CharacterLimbs.GetTile(1, 1),
								Rotation:       0,
								OffsetX:        0,
								OffsetY:        7,
								AnchorX:        1,
								AnchorY:        1,
								Darken:         0.02,
								ChildrenBefore: []world.CharacterFeature{
									{
										Texture:        textures.CharacterFeet.GetTile(1, 0),
										Rotation:       0,
										OffsetX:        0,
										OffsetY:        8,
										AnchorX:        1,
										AnchorY:        3,
										ChildrenBefore: nil,
										Darken:         0.02,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	player := &world.Character{
		Mask: world.CharacterMask{16, 28},
		Pos:  world.Coordinates{world.ChunkPixelSize * 1000, world.ChunkPixelSize * 105},
		Vel:  world.Vector{},
		Appearance: &world.CharacterAppearance{
			Body: alysia,
		},
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
