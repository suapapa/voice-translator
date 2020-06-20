package main

import (
	"image/color"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Game implements ebiten.Game interface.
type Game struct {
	Lang  string
	Str   string
	Start time.Time

	sync.RWMutex
}

var (
	game         Game
	screenWidth  = 1204
	screenHeight = 320
	maskColor    = color.NRGBA{0x00, 0x00, 0x00, 0xff}
)

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.Lock()
	defer g.Unlock()
	elapse := time.Now().Sub(g.Start)
	if elapse > 3*time.Second {
		g.Str = ""
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(maskColor)
	g.RLock()
	defer g.RUnlock()
	if g.Lang == "en" {
		drawTextBullon(screen, g.Str, color.White)
	} else {
		drawTextBullon(screen, g.Str, color.Black)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func gui() {
	// game := &Game{}
	game.Str = "Hello World"
	game.Start = time.Now()
	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("voice-translator")
	ebiten.SetMaxTPS(5)
	ebiten.SetRunnableOnUnfocused(true)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
