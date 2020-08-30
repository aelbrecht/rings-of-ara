package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"rings-of-ara/internal/draw"
	"rings-of-ara/internal/world"
)

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

	curX, curY := ebiten.CursorPosition()
	curWorldPos := g.World.Camera.ToWorld(world.Coordinates{int64(curX), int64(curY)})
	curChunkPos := curWorldPos.ToChunkPosition()
	curBlockPos := curWorldPos.ToBlockPosition()
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
		fmt.Sprintf("TPS: %0.2f\npos: %d,%d vel:%f,%f\n%s\n%s",
			ebiten.CurrentTPS(),
			g.World.Player.Pos.X,
			g.World.Player.Pos.Y,
			g.World.Player.Vel.X,
			g.World.Player.Vel.Y,
			chunkDebug,
			mouseDebug,
		),
	)
}
