package main

import "github.com/gen2brain/raylib-go/raylib"

func input() {

}

func update() {

}

func render() {

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.EndDrawing()
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		input()
		update()
		render()
	}

	rl.CloseWindow()
}
