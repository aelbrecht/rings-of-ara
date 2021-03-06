package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image"
	"image/color"
	"rings-of-ara/internal/draw"
	"rings-of-ara/internal/textures"
	"rings-of-ara/internal/world"
)

// draw loop
// uses a buffer to make drawing cleaner from different routines
func (g *Game) Draw(screen *ebiten.Image) {

	// fill sky
	//_ = screen.Fill(color.RGBA{228, 241, 254, 255})
	_ = screen.Fill(color.RGBA{
		R: 200,
		G: 220,
		B: 180,
		A: 255,
	})

	// fill background
	g.Buffers.BackgroundLayer.Clear()
	draw.BackgroundLayer(g.World, g.Buffers.BackgroundLayer)
	screen.DrawImage(g.Buffers.BackgroundLayer, nil)

	// fill block layer
	g.Buffers.BlockBackLayer.Clear()
	g.Buffers.BlockFrontLayer.Clear()
	draw.BlockLayer(g.World, g.Buffers.BlockFrontLayer, g.Buffers.BlockBackLayer)

	_ = screen.DrawImage(g.Buffers.BlockBackLayer, nil)

	draw.Character(g.World.Player, g.World, screen)

	_ = screen.DrawImage(g.Buffers.BlockFrontLayer, nil)

	chunkDebug := ""
	activeChunks := g.World.Camera.VisibleChunks()
	for _, chunk := range activeChunks {
		chunkDebug += fmt.Sprintf("%d,%d ", chunk.X, chunk.Y)
	}

	curX, curY := ebiten.CursorPosition()
	curWorldPos := g.World.Camera.ToWorld(world.Coordinates{int64(curX), int64(curY)})
	curBlockPos := curWorldPos.ToBlockPosition()
	curBlockCoordinates := curBlockPos.ToCoordinates()
	hightlightX, hightlightY := g.World.Camera.ToScreen(curBlockCoordinates).ValuesFloat()

	ops := ebiten.DrawImageOptions{}
	ops.GeoM.Translate(0, -10)
	ops.GeoM.Scale(3, 3)
	ops.GeoM.Translate(hightlightX, hightlightY)
	screen.DrawImage(textures.TileSetWorldInterface.SubImage(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{10, 10},
	}).(*ebiten.Image), &ops)

	curChunkPos := curWorldPos.ToChunkPosition()
	curRelBlockPos := curWorldPos.ToRelativeBlockPosition()
	curBlock := g.World.Planet.GetBlock(curWorldPos)
	var k uint16
	if curBlock != nil {
		k = curBlock.Kind
	}
	mouseDebug := fmt.Sprintf("s:%d,%d w:%d,%d c:%d,%d b%d,%d rb:%d,%d b.k:%d",
		curX, curY,
		curWorldPos.X, curWorldPos.Y,
		curChunkPos.X, curChunkPos.Y,
		curBlockPos.X, curBlockPos.Y,
		curRelBlockPos.X, curRelBlockPos.Y,
		k,
	)

	_ = ebitenutil.DebugPrint(screen,
		fmt.Sprintf("TPS: %0.0f\nFPS: %0.0f\nrender: %d\npos: %d,%d vel:%f,%f\n%s\n%s",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			draw.BlockRenders,
			g.World.Player.Pos.X,
			g.World.Player.Pos.Y,
			g.World.Player.Vel.X,
			g.World.Player.Vel.Y,
			chunkDebug,
			mouseDebug,
		),
	)
}
