package engine

import (
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type GameLoop func(delta float64)

type Game struct {
	Window   *pixelgl.Window
	gameLoop GameLoop
}

func startGame(g *Game) func() {
	return func() {
		cfg := pixelgl.WindowConfig{
			Title:  "Hello, Pixel!",
			Bounds: pixel.R(0, 0, 1024, 768),
			VSync:  true,
		}

		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			log.Fatal("Failed to create the window", err)
		}
		log.Println("Window created")

		win.SetSmooth(true)
		win.Clear(pixel.RGB(255, 255, 255))

		g.Window = win

		log.Println("Game started")

		frameStartTime := time.Now()

		for win.Closed() == false {
			delta := time.Since(frameStartTime).Seconds()
			frameStartTime = time.Now()

			g.gameLoop(delta)

			g.Window.Update()
		}
	}
}

func CreateGame() *Game {
	return &Game{}
}

func (g *Game) Run(GameLoop GameLoop) {
	g.gameLoop = GameLoop
	pixelgl.Run(startGame(g))
}
