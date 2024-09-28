package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(800, 600, "Snakey")

	log.Print("Hello Snakey!")
	defer func() {
		log.Print("Bye Snakey!")
		rl.CloseWindow()
	}()

	rl.SetTargetFPS(60)

	rl.DrawText("Snakey", 10, 10, 20, rl.White)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}
