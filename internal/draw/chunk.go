package draw

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
	"rings-of-ara/internal/world"
	"sync"
	"time"
)

var enqueuedChunks []world.ChunkPosition
var enqueuedChunksLock sync.Mutex
var enqueuedChunksSize = 0
var enqueuedCurrentCurrent = world.ChunkPosition{
	X: math.MaxInt32,
	Y: math.MaxInt32,
}

func init() {
	enqueuedChunks = make([]world.ChunkPosition, 200)
}

func Rectangle(size int, c color.RGBA, dst []byte) {
	for i := 0; i < size; i++ {
		SetPixel(i, 0, size, c, dst)
		SetPixel(0, i, size, c, dst)
		SetPixel(i, size-1, size, c, dst)
		SetPixel(size-1, i, size, c, dst)
	}
}

func enqueueChunk(coords world.ChunkPosition) {
	if enqueuedCurrentCurrent == coords {
		return
	}
	for i := 0; i < enqueuedChunksSize; i++ {
		if enqueuedChunks[i] == coords {
			return
		}
	}
	enqueuedChunksLock.Lock()
	chunkBufferLock.Lock()
	chunkBuffer[coords], _ = ebiten.NewImage(world.ChunkPointSize, world.ChunkPointSize, ebiten.FilterDefault)
	chunkBufferLock.Unlock()
	enqueuedChunks[enqueuedChunksSize] = coords
	enqueuedChunksSize += 1
	enqueuedChunksLock.Unlock()
}

func ClearChunks() {
	chunkBufferLock.Lock()
	chunkBuffer = make(map[world.ChunkPosition]*ebiten.Image)
	chunkBufferLock.Unlock()
}

func ChunkRenderer(w *world.Model) {

	for {
		time.Sleep(10 * time.Millisecond)

		if enqueuedChunksSize == 0 {
			continue
		}

		enqueuedChunksLock.Lock() // ENTER LOCK

		// get new active coordinates
		enqueuedCurrentCurrent = enqueuedChunks[enqueuedChunksSize-1]

		// get chunk if it has been generated
		w.Planet.Lock.Lock()
		ch := w.Planet.Chunks[enqueuedCurrentCurrent]
		w.Planet.Lock.Unlock()
		if ch == nil {
			enqueuedChunksLock.Unlock() // EXIT LOCK
			continue
		}

		// update queue
		enqueuedChunksSize -= 1

		enqueuedChunksLock.Unlock() // EXIT LOCK

		startTime := time.Now()

		chunkBufferLock.Lock()
		Chunk(ch,enqueuedCurrentCurrent, chunkBuffer[enqueuedCurrentCurrent], w.Debug)
		chunkBufferLock.Unlock()

		elapsed := time.Now().Sub(startTime)

		fmt.Printf("chunk %d,%d rendered in %dms \n", enqueuedCurrentCurrent.X, enqueuedCurrentCurrent.Y, elapsed.Milliseconds())
	}
}
