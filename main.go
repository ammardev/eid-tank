package main

import (
	"github.com/ammardev/eid-tank/internal/engine"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	game *engine.Game
	tank Tank
)

func gameLoop(delta float64) {
	for _, bullet := range bullets {
		bullet.Move(delta)
	}

	if game.Window.Pressed(pixelgl.KeyUp) {
		tank.Accelerate(delta)
	}

	if game.Window.Pressed(pixelgl.KeyRight) {
		tank.SteerRight(delta)
	}

	if game.Window.Pressed(pixelgl.KeyLeft) {
		tank.SteerLeft(delta)
	}

	if game.Window.Pressed(pixelgl.KeyDown) {
		tank.Reverse(delta)
	}

	if game.Window.JustPressed(pixelgl.KeySpace) {
		tank.Shoot()
	}

	game.Window.Clear(pixel.RGB(255, 255, 255))

	tank.Draw()

	for _, bullet := range bullets {
		bullet.Draw()
	}
}

func main() {
	game = engine.CreateGame()

	tank = createTank()

	game.Run(gameLoop)
}
