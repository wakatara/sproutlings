package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1000
	screenHeight = 480
	fps          = 60
)

var (
	running  = true
	bkgColor = rl.NewColor(147, 211, 196, 255)
)

func drawScene() {}

func input() {

}

func update() {
	running = !rl.WindowShouldClose()

}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)

	drawScene()

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sproutlings!")
	rl.SetExitKey(0)
	rl.SetTargetFPS(fps)
}

func quit() {
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}

	quit()
}
